package dayone2md

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func makeFileCataloger(dstdir string) *fileCataloger {
	return &fileCataloger{
		dstdir: dstdir,
	}
}

type fileInfo struct {
	Size int
	Hash string
}

// fileCataloger maintains a list of files that alread exist at the destination
// and files that have been generated to identify orphaned files.
type fileCataloger struct {
	dstdir string
	dst    map[string]fileInfo
	src    []string
}

// AddGeneratedFile note a generated file, relative to the output directory
func (z *fileCataloger) AddGeneratedFile(relname string) {
	slog.Debug("add src", "file", relname)
	z.src = append(z.src, relname)
}

// CatalogDestinationFiles saves each file in the given directory with the SHA hash of its contents.
func (z *fileCataloger) CatalogDestinationFiles(ctx context.Context) error {
	z.dst = make(map[string]fileInfo)

	visit := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("cannot load source file (%s): %w", path, err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if !f.IsDir() && f.Size() > 0 {
				hash, err := hashFile(path)
				if err != nil {
					if f.Mode()&os.ModeSymlink != 0 {
						// ignore symbolic links for files that do not exist
						return nil
					}
					return fmt.Errorf("cannot hash file (%s): %w", path, err)
				}
				relpath := strings.TrimPrefix(path, z.dstdir)
				slog.Debug("catalog file", "file", relpath, "hash", hash)
				z.dst[relpath] = fileInfo{
					Size: int(f.Size()),
					Hash: hash,
				}
			}
			return nil
		}
	}

	// HACK ensure trailing slash so that directories that are symlinks will be cataloged
	if !strings.HasSuffix(z.dstdir, "/") {
		z.dstdir += "/"
	}
	return filepath.Walk(z.dstdir, visit)
}

func (z *fileCataloger) ShouldWrite(filename string, payload []byte) bool {
	fileInfo, ok := z.dst[filename]
	if !ok {
		return true
	}
	if fileInfo.Size != len(payload) {
		return true
	}
	payloadHash, err := hash(bytes.NewReader(payload))
	if err != nil {
		// just log the error as a warning and allow copy
		slog.Warn("cannot hash payload", "err", err)
		return true
	}
	slog.Debug("shouldWrite", "file", filename, "hash", fileInfo.Hash, "payloadHash", payloadHash)
	return fileInfo.Hash == "" || payloadHash == "" || fileInfo.Hash != payloadHash
}

func (z *fileCataloger) ShouldWritePhoto(filename, dstFilename string, pg photoGetter) bool {
	fileInfo, ok := z.dst[dstFilename]
	if !ok {
		return true
	}
	payload, err := pg.GetPhoto(filename)
	if err != nil {
		slog.Warn("cannot get photo", "photo", filename, "err", err)
		return true
	}
	if fileInfo.Size != len(payload) {
		return true
	}
	payloadHash, err := hash(bytes.NewReader(payload))
	if err != nil {
		slog.Warn("cannot hash photo", "file", filename, "err", err)
		return true
	}
	slog.Debug("shouldWritePhoto", "file", filename, "hash", fileInfo.Hash, "payloadHash", payloadHash)
	return fileInfo.Hash == "" || payloadHash == "" || fileInfo.Hash != payloadHash
}

func (z *fileCataloger) RemoveOrphans() error {
	for targetFile := range z.dst {
		if !slices.Contains(z.src, targetFile) {
			absFilename := filepath.Join(z.dstdir, targetFile)
			if err := os.Remove(absFilename); err != nil {
				return fmt.Errorf("cannot remove orphan (%s): %w", targetFile, err)
			}
			slog.Info("removed orphan", "file", targetFile)
		}
	}
	return nil
}

func hashFile(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", nil
		}
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()
	return hash(f)
}

func hash(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
