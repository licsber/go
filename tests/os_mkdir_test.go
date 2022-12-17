package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestOsMkdir(t *testing.T) {
	err := os.Mkdir("debug", 0755)
	fmt.Println(err)
	err = os.Mkdir("debug", 0755)
	// 这样判断err是否为文件夹已存在 注意IsNotExist不是简单的取反
	if os.IsExist(err) {
		fmt.Println("文件夹已存在.")
	}
}
