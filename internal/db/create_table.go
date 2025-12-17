package db

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/irfan378/MiniDB/internal/storage"
)

const (
	CatalogPageNumber = 2

	CatalogHeaderSize = 4
	TableNameSize     = 32
	TableEntrySize    = 36
	MaxTablesPerPage  = (storage.PageSize - CatalogHeaderSize) / TableEntrySize
)

func (db *Database) CreateTable(name string) error {
	if len(name) == 0 || len(name) > TableNameSize {
		return errors.New("invalid table name length")
	}

	catalog, err := db.Pager.ReadPage(CatalogPageNumber)

	if err != nil {
		return err
	}
	count := binary.LittleEndian.Uint32(catalog[0:4])
	if count >= MaxTablesPerPage {
		return errors.New("catalog is full")
	}

	for i := uint32(0); i < count; i++ {
		entryOffset := CatalogHeaderSize + i*TableEntrySize
		nameBytes := catalog[entryOffset : entryOffset+TableNameSize]

		existingName := string(bytes.TrimRight(nameBytes, "\x00"))
		if existingName == name {
			return errors.New("table already exists")
		}
	}
	rootPage, err := db.Pager.AllocatePage()
	if err != nil {
		return err
	}

	entryOffset := CatalogHeaderSize + count*TableEntrySize
	copy(catalog[entryOffset:entryOffset+TableNameSize], []byte(name))
	binary.LittleEndian.PutUint32(
		catalog[entryOffset+TableNameSize:entryOffset+TableEntrySize],
		rootPage,
	)

	binary.LittleEndian.PutUint32(catalog[0:4], count+1)

	return db.Pager.WritePage(CatalogPageNumber, catalog)

}
