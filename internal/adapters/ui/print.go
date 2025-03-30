package ui

import (
	"fmt"
	"time"

	"github.com/AndresFWilT/afwtls/internal/domain/entities"
)

func PrintList(files []entities.File, nRecords int) {
	for _, file := range files[:nRecords] {
		style := entities.MapStyleByFileType[file.FileType]

		fmt.Printf("%11s %-8s %-8s %10d %s %s %s%s %s\n",
			file.Mode, file.UserName, file.GroupName, file.Size,
			file.ModificationTime.Format(time.DateTime), style.Icon,
			SetColor(file.Name, style.Color), style.Symbol, markHidden(file.IsHidden))
	}
}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}
	return Yellow("ø")
}
