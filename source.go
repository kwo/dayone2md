package dayone2md

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type source interface {
	journalGetter
	photoGetter
}

type journalGetter interface {
	GetJournal(context.Context, string) (*Journal, error)
}

type photoGetter interface {
	GetPhoto(string) ([]byte, error)
}

func newSource(location string) source {
	if strings.HasSuffix(strings.ToLower(location), ".zip") {
		return newZipSource(location)
	}
	if strings.HasSuffix(strings.ToLower(location), ".sqlite") {
		return newDBSource(location)
	}
	return newDirSource(filepath.Clean(location))
}

func newDirSource(dir string) *dirSource {
	return &dirSource{
		dir: dir,
	}
}

type dirSource struct {
	dir string
}

func (z *dirSource) GetJournal(_ context.Context, name string) (*Journal, error) {
	journalName := fmt.Sprintf("%s.json", name)
	data, err := z.getFile(journalName)
	if err != nil {
		return nil, fmt.Errorf("cannot find journal file: %s: %w", journalName, err)
	}
	j := &Journal{}
	if err := json.Unmarshal(data, j); err != nil {
		return nil, fmt.Errorf("cannot unmarshal journal json data: %w", err)
	}
	return j, nil
}

func (z *dirSource) GetPhoto(name string) ([]byte, error) {
	photoName := filepath.Join("photos", name)
	return z.getFile(photoName)
}

func (z *dirSource) getFile(name string) ([]byte, error) {
	fullpath := filepath.Join(z.dir, name)
	return os.ReadFile(fullpath)
}

func newZipSource(file string) *zipSource {
	return &zipSource{
		file: file,
	}
}

type zipSource struct {
	file string
}

func (z *zipSource) GetJournal(_ context.Context, name string) (*Journal, error) {
	journalName := fmt.Sprintf("%s.json", name)
	data, err := z.getFile(journalName)
	if err != nil {
		return nil, fmt.Errorf("cannot find journal zip entry: %s: %w", journalName, err)
	}
	j := &Journal{}
	if err := json.Unmarshal(data, j); err != nil {
		return nil, fmt.Errorf("cannot unmarshal journal json data: %w", err)
	}
	return j, nil
}

func (z *zipSource) GetPhoto(name string) ([]byte, error) {
	photoName := filepath.Join("photos", name)
	return z.getFile(photoName)
}

func (z *zipSource) getFile(name string) ([]byte, error) {
	r, err := zip.OpenReader(z.file)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()
	for _, f := range r.File {
		if f.Name == name {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer func() {
				_ = rc.Close()
			}()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("file not found: %s", name)
}
