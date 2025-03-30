package files

import (
	"path"
	"runtime"
	"strings"

	"github.com/AndresFWilT/afwtls/internal/domain/constants"
	"github.com/AndresFWilT/afwtls/internal/domain/entities"
	"github.com/AndresFWilT/afwtls/internal/usecase/files/info"
)

func SetFileType(f *entities.File) {
	switch {
	case isLink(*f):
		f.FileType = constants.FileLink
	case f.IsDir:
		f.FileType = constants.FileDirectory
	case IsExec(*f):
		f.FileType = constants.FileExecutable
	case isCompress(*f):
		f.FileType = constants.FileCompress
	case isImage(*f):
		f.FileType = constants.FileImage
	default:
		f.FileType = constants.FileRegular
	}
}

func IsHidden(fileName, basePath string) bool {
	filePath := fileName

	if runtime.GOOS == constants.Windows {
		filePath = path.Join(basePath, filePath)
	}

	return info.IsHidden(filePath)
}

func isImage(f entities.File) bool {
	return strings.HasSuffix(f.Name, constants.Png) || strings.HasSuffix(f.Name, constants.Jpg) ||
		strings.HasSuffix(f.Name, constants.Gif)
}

func isCompress(f entities.File) bool {
	return strings.HasSuffix(f.Name, constants.Zip) || strings.HasSuffix(f.Name, constants.Gz) ||
		strings.HasSuffix(f.Name, constants.Tar) || strings.HasSuffix(f.Name, constants.Rar) ||
		strings.HasSuffix(f.Name, constants.Deb)
}

func IsExec(file entities.File) bool {
	if runtime.GOOS == constants.Windows {
		return strings.HasSuffix(file.Name, constants.Exe)
	}

	return strings.Contains(file.Mode, "x")
}

func isLink(file entities.File) bool {
	return strings.HasPrefix(strings.ToLower(file.Mode), "l")
}
