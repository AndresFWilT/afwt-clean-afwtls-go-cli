package files

import (
	"sort"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/AndresFWilT/afwtls/internal/domain/entities"
)

func OrderByName(files []entities.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			strings.ToLower(files[i].Name),
			strings.ToLower(files[j].Name),
			isReverse,
		)
	})
}

func OrderBySize(files []entities.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].Size,
			files[j].Size,
			isReverse,
		)
	})
}

func OrderByTime(files []entities.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].ModificationTime.Unix(),
			files[j].ModificationTime.Unix(),
			isReverse,
		)
	})
}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}

	return i < j
}
