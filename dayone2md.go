package dayone2md

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"
	"text/template"
	"time"
)

type Options interface {
	GetJournalName() string
	GetInputLocation() string
	GetOutputDirectory() string
	GetTemplate() string
	IsGroupByDay() bool
	IsSortReverse() bool
	IsRemoveOrphans() bool
}

// Convert export DayOne entries to markdown.
func Convert(ctx context.Context, opts Options) error {
	src := newSource(opts.GetInputLocation())

	j, err := src.GetJournal(opts.GetJournalName())
	if err != nil {
		return fmt.Errorf("cannot load journal: %w", err)
	}

	if err := ensureDirectories(opts.GetOutputDirectory()); err != nil {
		return err
	}

	tpl, err := loadTemplate(opts.GetTemplate())
	if err != nil {
		return fmt.Errorf("cannot load template (%s): %w", opts.GetTemplate(), err)
	}

	catalog := makeFileCataloger(opts.GetOutputDirectory())
	if err := catalog.CatalogDestinationFiles(ctx); err != nil {
		return fmt.Errorf("cannot catalog files: %w", err)
	}

	// the keys are separated out because we like to write files out by date
	entries, filenamesByFilename, filenames := sortEntries(j, opts.IsGroupByDay(), opts.IsSortReverse())

	for i := 0; i < len(j.Entries); i++ {
		e := &j.Entries[i]
		// sort the tags, as not always the same from json import
		sort.Strings(e.Tags)
		// remove escape characters
		e.Text = strings.ReplaceAll(e.Text, `\`, "")

		// save photos
		photos, err := savePhotos(e.Photos, opts.GetOutputDirectory(), src, catalog)
		if err != nil {
			return err
		}
		e.Text = replacePhotoLinks(e.Text, photos)
		e.Text = wikilinks(e.Text, filenamesByFilename)

		// separate title from body
		body, title := trimText(e.Text)
		e.Title = prepareTitle(title)
		e.Text = body
	}

	for _, filename := range filenames {
		dayEntries := entries[filename]

		// slog.Debug("entry", "date", dt)
		entryFilenameRel := fmt.Sprintf("%s.md", filename)
		entryFilename := filepath.Join(opts.GetOutputDirectory(), entryFilenameRel)

		generatedText, err := generateText(dayEntries, tpl)
		if err != nil {
			return err
		}

		payload := []byte(generatedText)
		if catalog.ShouldWrite(entryFilenameRel, payload) {
			slog.Info("writing file", "file", entryFilename)
			if err := os.WriteFile(entryFilename, payload, 0o600); err != nil {
				return err
			}
		}
		catalog.AddGeneratedFile(entryFilenameRel)
	} // entries loop

	// remove orphans
	if opts.IsRemoveOrphans() {
		if err := catalog.RemoveOrphans(); err != nil {
			return fmt.Errorf("cannot remove orphans: %w", err)
		}
	}

	return nil
}

func ensureDirectories(outputDir string) error {
	// create directory, if not exist
	if stats, err := os.Stat(outputDir); err != nil {
		// create dir
		if err := os.MkdirAll(outputDir, 0o755); err != nil {
			return fmt.Errorf("cannot create output directory: %s: %w", outputDir, err)
		}
	} else if !stats.IsDir() {
		return fmt.Errorf("cannot create output directory: %s, because the name already exists as a non-directory", outputDir)
	}

	// create photos subdirectory, if not exist
	photosDirectory := filepath.Join(outputDir, "photos")
	if stats, err := os.Stat(photosDirectory); err != nil {
		// create dir
		if err := os.Mkdir(photosDirectory, 0o755); err != nil {
			return fmt.Errorf("cannot create photos directory: %s: %w", photosDirectory, err)
		}
	} else if !stats.IsDir() {
		return fmt.Errorf("cannot create photos directory: %s, because the name already exists as a non-directory", photosDirectory)
	}

	return nil
}

// sortEntries returns maps of Entry keyed by filename, filename keyed by UUID, and the filenames sorted, ascending
func sortEntries(j *Journal, groupByDay, sortWithinDaysReverse bool) (map[string][]*Entry, map[string]string, []string) {
	// assign Date attr to all entries
	tzCache := makeTimeZoneCache()
	for i := 0; i < len(j.Entries); i++ {
		tzLocation := tzCache(j.Entries[i].TimeZone)
		d1 := j.Entries[i].CreationDate.In(tzLocation)
		j.Entries[i].Date = d1
	}

	// group entries by date or date-time
	entries := make(map[string][]*Entry)
	filenames := make(map[string]string)
	for i := 0; i < len(j.Entries); i++ {
		e := &j.Entries[i]
		filename := calcFilename(e.Date, groupByDay)
		entries[filename] = append(entries[filename], e)
		filenames[e.UUID] = filename
	}

	// sort day entries
	for _, dayEntries := range entries {
		slices.SortFunc(dayEntries, func(a, b *Entry) int {
			if sortWithinDaysReverse {
				return b.CreationDate.Compare(a.CreationDate)
			}
			return a.CreationDate.Compare(b.CreationDate)
		})
	}

	// sort filenames
	names := make([]string, 0, len(entries))
	for k := range entries {
		names = append(names, k)
	}
	slices.Sort(names)

	return entries, filenames, names
}

func calcFilename(dt time.Time, groupByDay bool) string {
	if groupByDay {
		return dt.Format("2006-01-02")
	}
	return dt.Format("20060102T1504")
}

func generateText(entries []*Entry, tmpl *template.Template) (string, error) {
	w := &bytes.Buffer{}
	if err := tmpl.Execute(w, entries); err != nil {
		return "", err
	}
	text := w.String()
	text = strings.TrimSpace(text) + "\n" // trim all leading/trailing newlines/whitespace and finish with a single newline
	return text, nil
}

func trimText(text string) (string, string) {
	text = strings.TrimSpace(text) + "\n"  // trim all leading/trailing newlines/whitespace and finish with a single newline
	lines := strings.SplitN(text, "\n", 2) // because we add newline above, len(lines) will always be 2
	if lines[1] == "" {
		return text, ""
	}
	title := strings.TrimSpace(lines[0])
	body := strings.Join(lines[1:], "\n")
	body = strings.TrimSpace(body) + "\n"
	return body, title
}

func prepareTitle(title string) string {
	title = strings.TrimLeft(title, "# ") // remove markdown
	if strings.HasPrefix(title, "2") {    // if the title is a date, use an alternate title
		title = ""
	}
	return title
}

func wikilinks(text string, filenamesByUUID map[string]string) string {
	// replace dayone link references with wikilinks
	re := regexp.MustCompile(`\[([\w\s]+)\]\(dayone\:\/\/view\?entryId\=([A-Z0-9]+)\)`)
	submatches := re.FindAllStringSubmatch(text, -1)
	for _, submatch := range submatches {
		subtext, title, uuid := submatch[0], submatch[1], submatch[2]
		if filename, ok := filenamesByUUID[uuid]; ok {
			wikilink := fmt.Sprintf("[[%s|%s]]", filename, title)
			text = strings.Replace(text, subtext, wikilink, 1)
		}
	}
	return text
}
