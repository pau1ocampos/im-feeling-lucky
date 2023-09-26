package lucky

import (
	"encoding/json"
	"io"
	"os"
)

type HdFileSystem struct{}

func (o *Options) ParseFromFile(path string) (*Draws, error) {
	b, err := o.FileSystem.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var draws *Draws
	err = json.Unmarshal(b, &draws)
	if err != nil {
		return nil, err
	}

	return draws, err
}

func (fs HdFileSystem) ReadFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}
