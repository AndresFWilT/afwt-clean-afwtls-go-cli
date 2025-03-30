package main

import (
	flagAdapter "github.com/AndresFWilT/afwtls/internal/adapters/flag"
	"github.com/AndresFWilT/afwtls/internal/adapters/ui"
	"github.com/AndresFWilT/afwtls/internal/usecase/files"
	flagCommand "github.com/AndresFWilT/afwtls/internal/usecase/flag"
)

func main() {
	flagAdapter.ParseFlag()

	path := flagAdapter.SetArg(0)
	if path == "" {
		path = "."
	}

	filesList := files.NewList(path)
	fs, err := filesList.Execute()

	if err != nil {
		panic(err)
	}

	fc := &flagCommand.Command{
		HasOrderReverse: *flagAdapter.HasOrderReverse,
		HasOrderBySize:  *flagAdapter.HasOrderBySize,
		HasOrderByTime:  *flagAdapter.HasOrderByTime,
		NumberRecords:   *flagAdapter.NumberRecords,
	}
	fc.Execute(fs)

	ui.PrintList(fs, fc.NumberRecords)
}
