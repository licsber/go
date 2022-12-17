package lFile

import (
	"fmt"
	"strconv"
)

var FriendlySizeRound = strconv.Itoa(2)

func FriendlySize(size int64) string {
	s := float64(size)

	if size < 0 {
		fmt.Print("文件大小非法: ")
		fmt.Println(size)
		return "0.00B"
	} else if size < KB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fB", s/float64(B))
	} else if size < MB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fKB", s/float64(KB))
	} else if size < GB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fMB", s/float64(MB))
	} else if size < TB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fGB", s/float64(GB))
	} else if size < PB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fTB", s/float64(TB))
	} else if size < EB {
		return fmt.Sprintf("%."+FriendlySizeRound+"fPB", s/float64(PB))
	} else {
		return fmt.Sprintf("%."+FriendlySizeRound+"fEB", float64(size)/float64(EB))
	}
}
