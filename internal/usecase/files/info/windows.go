//go:build windows
// +build windows

package info

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func IsHidden(filename string) bool {
	pointer, err := windows.UTF16PtrFromString(filename)
	if err != nil {
		panic(fmt.Sprint(filename, err))
	}
	attributes, err := windows.GetFileAttributes(pointer)
	if err != nil {
		panic(fmt.Sprint(filename, err))
	}
	return attributes&windows.FILE_ATTRIBUTE_HIDDEN != 0
}

func GetUserAndGroup(infoSys any) (userName, groupName string) {
	return
}
