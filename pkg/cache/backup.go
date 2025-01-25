package cache

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/fluffy-melli/MoliDB/pkg/gzip"
)

const BACKUP = "./backup/"

func FileList() ([]string, error) {
	files, err := os.ReadDir(BACKUP)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, file := range files {
		list = append(list, file.Name())
	}
	return list, nil
}

func LastLoad() *Safe {
	list, err := FileList()
	if err != nil {
		return NewSafeMap()
	}
	if len(list) == 0 {
		return NewSafeMap()
	}
	load, err := Load(list[len(list)-1])
	if err != nil {
		return NewSafeMap()
	}
	return load
}

func Load(filename string) (*Safe, error) {
	safe := NewSafeMap()
	file, err := os.Open(filepath.Join(BACKUP, filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	safe.store, err = gzip.Decompress(data)
	if err != nil {
		return nil, err
	}
	return safe, nil
}

func (sm *Safe) Backup() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if err := os.MkdirAll(BACKUP, os.ModePerm); err != nil {
		return err
	}
	data, err := gzip.Compress(sm.store)
	if err != nil {
		return err
	}
	filename := time.Now().Format("2006-01-02T15-04-05.bak")
	file, err := os.Create(filepath.Join(BACKUP, filename))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
