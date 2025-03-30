package flag

import "flag"

var (
	Pattern       = flag.String("p", "", "filter by pattern")
	All           = flag.Bool("a", false, "all files including hidden files")
	NumberRecords = flag.Int("n", 0, "number of records")

	HasOrderByTime  = flag.Bool("t", false, "sort by time, oldest first")
	HasOrderBySize  = flag.Bool("s", false, "sort by file size, smallest first")
	HasOrderReverse = flag.Bool("r", false, "reverse order while sorting")
)

func ParseFlag() {
	flag.Parse()
}

func SetArg(arg int) string {
	return flag.Arg(arg)
}
