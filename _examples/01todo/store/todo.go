package store

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Store interface {
	Load(ctx context.Context, name string, data interface{}) error
	Save(ctx context.Context, name string, data interface{}) error
}

type FileStore struct {
	Dir string
}

func (s *FileStore) Load(ctx context.Context, name string, data interface{}) error {
	filename := filepath.Join(s.Dir, name)
	f, err := os.Open(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return xerrors.Errorf("open file: %w", err)
		}

		log.Printf("%s is not found, create file", filename)
		if err := s.Save(ctx, name, data); err != nil {
			return xerrors.Errorf("create file (does not exist): %w", err)
		}

		f, err = os.Open(filename)
		if err != nil {
			return xerrors.Errorf("open file (after create): %w", err)
		}
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	if err := dec.Decode(data); err != nil {
		return xerrors.Errorf("decode : %w", err)
	}
	return nil
}

func (s *FileStore) Save(ctx context.Context, name string, data interface{}) error {
	filename := filepath.Join(s.Dir, name)
	f, err := os.Create(filename)
	if err != nil {
		return xerrors.Errorf("create file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return xerrors.Errorf("encode : %w", err)
	}
	return nil
}
