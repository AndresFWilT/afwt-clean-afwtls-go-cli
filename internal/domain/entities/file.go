package entities

import (
	"time"
	
	"github.com/fatih/color"

	"github.com/AndresFWilT/afwtls/internal/domain/constants"
)

type File struct {
	Name             string
	FileType         int
	IsDir            bool
	IsHidden         bool
	UserName         string
	GroupName        string
	Size             int64
	ModificationTime time.Time
	Mode             string
}

type styleFileType struct {
	Symbol string
	Color  color.Attribute
	Icon   string
}

var MapStyleByFileType = map[int]styleFileType{
	constants.FileRegular:    {Icon: "📄"},
	constants.FileDirectory:  {Icon: "📂", Color: color.FgBlue, Symbol: "/"},
	constants.FileExecutable: {Icon: "🚀", Color: color.FgGreen, Symbol: "*"},
	constants.FileCompress:   {Icon: "📦", Color: color.FgRed},
	constants.FileImage:      {Icon: "📸", Color: color.FgMagenta},
	constants.FileLink:       {Icon: "🔗", Color: color.FgCyan},
}
