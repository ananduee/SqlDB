package storage

import "os"

type pagerBlock struct {
	file     *os.File
	fileSize int64
	pages    [maxPagesInTable]*page
}

func newPagerBlock(filename string) (*pagerBlock, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	p := &pagerBlock{file: file, fileSize: fileInfo.Size()}
	return p, nil
}

func (pagerInstance *pagerBlock) isPageOnDisk(pageNumber uint) bool {
	rowsOnDisk := uint(pagerInstance.fileSize) / rowSize
	pagesOnDisk := rowsOnDisk / rowsPerPage
	return pageNumber < pagesOnDisk
}

func (pagerInstance *pagerBlock) getPage(pageNumber int) {
	if pageNumber > maxPagesInTable {
		panic("pageNumber > maxPagesInTable")
	}

	pageInMemory := pagerInstance.pages[pageNumber]
	if pageInMemory == nil {
		// Cache miss we need to load this page from disk
		// if available.
		if pagerInstance.isPageOnDisk(pageNumber) {
			// Read file from disk.

		}
	}
}
