package flag

import (
	"github.com/AndresFWilT/afwtls/internal/domain/entities"
	"github.com/AndresFWilT/afwtls/internal/usecase/files"
)

type Command struct {
	HasOrderBySize  bool
	HasOrderByTime  bool
	HasOrderReverse bool
	NumberRecords   int
}

func (c *Command) Execute(fs []entities.File) []entities.File {
	switch {
	case c.HasOrderBySize:
		files.OrderBySize(fs, c.HasOrderBySize)
	case c.HasOrderByTime:
		files.OrderByTime(fs, c.HasOrderByTime)
	default:
		files.OrderByName(fs, c.HasOrderReverse)
	}

	if c.NumberRecords == 0 || c.NumberRecords > len(fs) {
		c.NumberRecords = len(fs)
	}

	return fs
}
