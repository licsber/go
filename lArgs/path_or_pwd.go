package lArgs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
PathOrPWD

从全部命令行参数中拼接路径 默认$PWD.

feature-0: 返回绝对路径且保证路径存在.

feature-1: 支持直接粘贴中间带空格的文件名.
*/
func PathOrPWD() (p string, fi os.FileInfo) {
	if len(os.Args) != 1 {
		p = strings.Join(os.Args[1:], " ")
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			panic("check if $PWD exists?")
		}
		p = pwd
	}

	p, err := filepath.Abs(p)
	if err != nil {
		panic("check if $PWD exists?")
	}

	fi, err = os.Stat(p)
	if err != nil {
		panic(err)
	}

	fmt.Println("工作路径: " + p)
	return
}
