package dayone2md

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// savePhotos - TODO too complex.
func savePhotos(entryPhotos []Photo, outputDir string, pg photoGetter, catalog *fileCataloger) (map[string]string, error) {
	// save photos
	photos := make(map[string]string)
	for _, photo := range entryPhotos {
		slog.Debug("savePhotos", "photo.md5", photo.MD5, "photo.type", photo.Type, "photo.filename", photo.Filename)
		// find file in export photos directory
		sourceFile := fmt.Sprintf("%s.%s", photo.MD5, photo.Type)
		targetFile := photo.Filename
		if targetFile == "" {
			targetFile = sourceFile
		}
		targetFile = filepath.Join("photos", targetFile)
		photos[photo.Identifier] = targetFile
		// copy file from archive to destination
		targetPath := filepath.Join(outputDir, targetFile)
		slog.Debug("savePhotos2", "sourceFile", sourceFile, "targetFile", targetFile)
		if catalog.ShouldWritePhoto(sourceFile, targetFile, pg) {
			slog.Info("copying photo", "file", targetPath)
			if err := copyPhoto(sourceFile, targetPath, pg); err != nil {
				return nil, fmt.Errorf("cannot copy photo (%s): %w", targetFile, err)
			}
		}
		catalog.AddGeneratedFile(targetFile)
	}
	return photos, nil
}

func replacePhotoLinks(text string, photos map[string]string) string {
	// replace dayone photo references with markdown image references in text
	re := regexp.MustCompile(`dayone\-moment\:\/\/(\w+)`)
	submatches := re.FindAllStringSubmatch(text, -1)
	for _, submatch := range submatches {
		subtext, id := submatch[0], submatch[1]
		if filename, ok := photos[id]; ok {
			link := filepath.Join(".", filename)
			text = strings.Replace(text, subtext, link, 1)
		}
	}
	return text
}

func copyPhoto(source, destination string, pg photoGetter) error {
	srcData, err := pg.GetPhoto(source)
	if err != nil {
		return fmt.Errorf("cannot open photo source file: %w", err)
	}
	if err := os.WriteFile(destination, srcData, 0o600); err != nil {
		return fmt.Errorf("cannot copy photo: %w", err)
	}
	return nil
}
