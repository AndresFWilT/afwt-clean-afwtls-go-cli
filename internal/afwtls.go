package internal

import (
	"time"

	"github.com/fatih/color"
)

const Windows = "windows"

const (
	FileRegular int = iota
	FileDirectory
	FileExecutable
	FileCompress
	FileImage
	FileLink
)

const (
	Exe = ".exe"
	Deb = ".deb"
	Zip = ".zip"
	Gz  = ".gz"
	Tar = ".tar"
	Rar = ".rar"
	Png = ".png"
	Jpg = ".jpg"
	Gif = ".gif"
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
	FileRegular:    {Icon: "📄"},
	FileDirectory:  {Icon: "📂", Color: color.FgBlue, Symbol: "/"},
	FileExecutable: {Icon: "🚀", Color: color.FgGreen, Symbol: "*"},
	FileCompress:   {Icon: "📦", Color: color.FgRed},
	FileImage:      {Icon: "📸", Color: color.FgMagenta},
	FileLink:       {Icon: "🔗", Color: color.FgCyan},
}

var (
	Blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	Green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	Red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	Magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	Cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
)
