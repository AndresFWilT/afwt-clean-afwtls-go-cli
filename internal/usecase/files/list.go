package files

import (
	"fmt"
	"os"
	"regexp"

	"github.com/AndresFWilT/afwtls/internal/adapters/flag"
	"github.com/AndresFWilT/afwtls/internal/domain/entities"
)

type List struct {
	path string
}

func NewList(path string) *List {
	return &List{
		path: path,
	}
}

func (l *List) Execute() ([]entities.File, error) {
	dirs, err := os.ReadDir(l.path)
	if err != nil {
		return nil, fmt.Errorf("error trying to get directory path: %v", err)
	}
	var fs []entities.File
	for _, dir := range dirs {
		isHidden := IsHidden(dir.Name(), l.path)

		if isHidden && !*flag.All {
			continue
		}

		if *flag.Pattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flag.Pattern, dir.Name())
			if err != nil {
				panic(err)
			}

			if !isMatched {
				continue
			}
		}

		f, err := GetFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	return fs, nil
}
