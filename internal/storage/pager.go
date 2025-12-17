package storage

import (
	"fmt"
	"io"
	"os"
)

const PageSize = 4096

type Pager struct {
	file *os.File
}

func OpenPager(path string) (*Pager, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &Pager{file: file}, nil
}

func (p *Pager) ReadPage(pageNum uint32) ([]byte, error) {

	fmt.Println("error is here")

	fmt.Println(pageNum)
	fmt.Println(p.NumPages())
	if pageNum >= p.NumPages() {
		return nil, io.EOF
	}

	fmt.Println("Here after page")
	buf := make([]byte, PageSize)

	_, err := p.file.ReadAt(buf, int64(pageNum)*PageSize)
	return buf, err
}

func (p *Pager) WritePage(pageNum uint32, data []byte) error {
	offset := int64((pageNum) * PageSize)
	_, err := p.file.WriteAt(data, offset)
	return err
}

func (p *Pager) NumPages() uint32 {
	info, _ := p.file.Stat()

	return uint32(info.Size() / PageSize)
}

func (p *Pager) AllocatePage() (uint32, error) {
	pageNum := p.NumPages()

	empty := make([]byte, PageSize)
	_, err := p.file.WriteAt(empty, int64(pageNum)*PageSize)
	if err != nil {
		return 0, err
	}

	return pageNum, nil
}
