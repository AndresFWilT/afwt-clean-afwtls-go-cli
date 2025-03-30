package files

import (
	"io/fs"

	"github.com/AndresFWilT/afwtls/internal/domain/entities"
	fileinfo "github.com/AndresFWilT/afwtls/internal/usecase/files/info"
)

func GetFile(dir fs.DirEntry, isHidden bool) (entities.File, error) {
	info, err := dir.Info()
	if err != nil {
		return entities.File{}, err
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := entities.File{
		Name:             dir.Name(),
		IsDir:            dir.IsDir(),
		IsHidden:         isHidden,
		Size:             info.Size(),
		Mode:             info.Mode().String(),
		ModificationTime: info.ModTime(),
		UserName:         userName,
		GroupName:        groupName,
	}
	SetFileType(&f)

	return f, nil
}
