package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/AJRDRGZ/fileinfo"
	"github.com/AndresFWilT/afwtls/internal"
	"github.com/fatih/color"
	"golang.org/x/exp/constraints"
)

func main() {
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hide files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by file size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []internal.File{}
	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		if *flagPattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name())
			if err != nil {
				panic(err)
			}

			if !isMatched {
				continue
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	if !*hasOrderBySize || !*hasOrderByTime {
		orderByName(fs, *hasOrderReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderReverse)
	}

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}

	printList(fs, *flagNumberRecords)
}

func getFile(dir fs.DirEntry, isHidden bool) (internal.File, error) {
	info, err := dir.Info()
	if err != nil {
		return internal.File{}, err
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := internal.File{
		Name:             dir.Name(),
		IsDir:            dir.IsDir(),
		IsHidden:         isHidden,
		Size:             info.Size(),
		Mode:             info.Mode().String(),
		ModificationTime: info.ModTime(),
		UserName:         userName,
		GroupName:        groupName,
	}
	setFileType(&f)

	return f, nil
}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}

	return i < j
}

func orderByName(files []internal.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			strings.ToLower(files[i].Name),
			strings.ToLower(files[j].Name),
			isReverse,
		)
	})
}

func orderBySize(files []internal.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].Size,
			files[j].Size,
			isReverse,
		)
	})
}

func orderByTime(files []internal.File, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].ModificationTime.Unix(),
			files[j].ModificationTime.Unix(),
			isReverse,
		)
	})
}

func printList(files []internal.File, nRecords int) {
	for _, file := range files[:nRecords] {
		style := internal.MapStyleByFileType[file.FileType]

		fmt.Printf("%11s %-8s %-8s %10d %s %s %s%s %s\n",
			file.Mode, file.UserName, file.GroupName, file.Size,
			file.ModificationTime.Format(time.DateTime), style.Icon,
			setColor(file.Name, style.Color), style.Symbol, markHidden(file.IsHidden))
	}
}

func setFileType(f *internal.File) {
	switch {
	case isLink(*f):
		f.FileType = internal.FileLink
	case f.IsDir:
		f.FileType = internal.FileDirectory
	case IsExec(*f):
		f.FileType = internal.FileExecutable
	case isCompress(*f):
		f.FileType = internal.FileCompress
	case isImage(*f):
		f.FileType = internal.FileImage
	default:
		f.FileType = internal.FileRegular
	}
}

func isImage(f internal.File) bool {
	return strings.HasSuffix(f.Name, internal.Png) || strings.HasSuffix(f.Name, internal.Jpg) ||
		strings.HasSuffix(f.Name, internal.Gif)
}

func isCompress(f internal.File) bool {
	return strings.HasSuffix(f.Name, internal.Zip) || strings.HasSuffix(f.Name, internal.Gz) ||
		strings.HasSuffix(f.Name, internal.Tar) || strings.HasSuffix(f.Name, internal.Rar) ||
		strings.HasSuffix(f.Name, internal.Deb)
}

func IsExec(file internal.File) bool {
	if runtime.GOOS == internal.Windows {
		return strings.HasSuffix(file.Name, internal.Exe)
	}

	return strings.Contains(file.Mode, "x")
}

func isLink(file internal.File) bool {
	return strings.HasPrefix(strings.ToLower(file.Mode), "l")
}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return internal.Blue(nameFile)
	case color.FgGreen:
		return internal.Green(nameFile)
	case color.FgRed:
		return internal.Red(nameFile)
	case color.FgMagenta:
		return internal.Magenta(nameFile)
	case color.FgCyan:
		return internal.Cyan(nameFile)
	default:
		return nameFile
	}
}

func isHidden(fileName, basePath string) bool {
	filePath := fileName

	if runtime.GOOS == internal.Windows {
		filePath = path.Join(basePath, filePath)
	}

	return fileinfo.IsHidden(filePath)
}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}

	return internal.Yellow("ø")
}
