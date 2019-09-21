package storage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"unsafe"

	"github.com/ananduee/SqlDB/compiler"
)

const (
	pageSize        = 4096
	maxPagesInTable = 100
	rowSize         = uint(unsafe.Sizeof(compiler.Row{}))
	rowsPerPage     = pageSize / rowSize
)

var (
	ErrorTableFull = errors.New("Table is full can't insert")
)

type page [pageSize]byte

// MemoryTable is in memory block for storage.
type MemoryTable struct {
	rowsCount uint
	pages     [maxPagesInTable]*page
}

// NewMemoryTable Create new instance.
func NewMemoryTable() *MemoryTable {
	return &MemoryTable{rowsCount: 0}
}

func getLocationToInsert(rowsInTable uint) (pageNum, indexInPage uint) {
	pageNum = rowsInTable / rowsPerPage
	indexInPage = rowsInTable % rowsPerPage
	return
}

func insertIntoPage(p *page, row compiler.Row, indexInPage uint) {
	binaryR := new(bytes.Buffer)
	err := binary.Write(binaryR, binary.LittleEndian, row)
	if err != nil {
		panic("failed to marshal row")
	}
	if uint(len(binaryR.Bytes())) != rowSize {
		// struct has some extra padding!
		panic("size of binary representation different from struct size")
	}
	copy(p[indexInPage*rowSize:], binaryR.Bytes())
}

// Insert a new row in table
func (table *MemoryTable) Insert(row compiler.Row) error {
	pageNum, indexInPage := getLocationToInsert(table.rowsCount)
	if pageNum >= maxPagesInTable {
		return ErrorTableFull
	}
	pageInstance := table.pages[pageNum]
	if pageInstance == nil {
		pageInstance = &page{}
		table.pages[pageNum] = pageInstance
	}
	insertIntoPage(pageInstance, row, indexInPage)
	table.rowsCount++
	return nil
}

// GetRows returns all rows present in the table.
func (table *MemoryTable) GetRows() (rows []compiler.Row, err error) {
	for _, page := range table.pages {
		if page == nil {
			break
		}
		// read page
		bytesArray := [pageSize]byte(*page)
		row := &compiler.Row{}
		emptyRow := compiler.Row{}
		bytesReader := bytes.NewReader(bytesArray[:])
		var rowsRead uint = 0
		for {
			if rowsRead > rowsPerPage {
				break
			}
			err := binary.Read(bytesReader, binary.LittleEndian, row)
			// if there are no contents or it returns empty row means page
			// does not have any more rows and we have hit end.
			if err == io.ErrUnexpectedEOF || *row == emptyRow {
				break
			}
			if err != nil {
				panic(err)
			}
			rows = append(rows, *row)
			rowsRead++
		}
	}
	return rows, nil
}
