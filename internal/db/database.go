package db

import (
	"encoding/binary"

	"github.com/irfan378/MiniDB/internal/storage"
)

type Database struct {
	Pager *storage.Pager
}

func Open(path string) (*Database, error) {
	pager, err := storage.OpenPager(path)
	if err != nil {
		return nil, err
	}
	size := pager.NumPages()
	if size == 0 {
		initDatabase(pager)
	}

	return &Database{Pager: pager}, nil
}
func initDatabase(pager *storage.Pager) {
	header := make([]byte, storage.PageSize)
	copy(header[0:8], []byte("MINIDB1"))
	binary.LittleEndian.PutUint32(header[8:12], storage.PageSize)
	binary.LittleEndian.PutUint32(header[12:16], 2)
	pager.WritePage(1, header)
	catalog := make([]byte, storage.PageSize)
	binary.LittleEndian.PutUint32(catalog[0:4], 0)
	pager.WritePage(2, catalog)
}
