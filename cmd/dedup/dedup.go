package main

import (
	"fmt"
	"github.com/licsber/go/lArray"
	"github.com/licsber/go/lErr"
	"os"
	"path/filepath"
)

var Path2Sha256 = make(map[string]string)

func deDupFile(srcPath, cmpPath string) {
	// 从缓存里拿源文件Hash 拿不到就计算后放到缓存
	srcSha256, ok := Path2Sha256[srcPath]
	if !ok {
		srcSha256 = CalSha256Hex(srcPath)
		Path2Sha256[srcPath] = srcSha256
	}

	cmpSha256 := CalSha256Hex(cmpPath)
	if srcSha256 == cmpSha256 {
		fmt.Println()
		fmt.Print("Found same sha256: ")
		fmt.Println(srcSha256)
		fmt.Print("SRC: ")
		fmt.Println(srcPath)
		fmt.Print("DST: ")
		fmt.Println(cmpPath)
		if !*ForceFlag {
			fmt.Print("RM(DRY): ")
			fmt.Println(cmpPath)
			return
		}

		fmt.Print("RM: ")
		fmt.Println(cmpPath)
		err := os.Remove(cmpPath)
		lErr.PanicErr(err)
	}
}

func deDupFiles(srcPaths, cmpPaths []string) {
	// 用这种比较方法小心出现两个源文件size+headHash都相同
	// 但是Hash不同结果目标文件被跳过的情况
	srcHeadHashes2PossiblePaths := make(map[string][]string)

	// 首先计算下所有源文件的头Hash
	for _, path := range srcPaths {
		headSha256 := CalHeadSha256Hex(path)
		srcHeadHashes2PossiblePaths[headSha256] = append(srcHeadHashes2PossiblePaths[headSha256], path)
	}

	for _, cmpPath := range cmpPaths {
		// 计算所有目标文件的头Hash
		headSha256 := CalHeadSha256Hex(cmpPath)
		// 拿头Hash两边比较 不要耗时直接计算整体Hash
		possiblePaths, ok := srcHeadHashes2PossiblePaths[headSha256]
		if !ok {
			continue
		}

		// 跳过源和目标相同的情况 避免一份副本都不剩下
		if lArray.StringArrayContains(possiblePaths, cmpPath) {
			continue
		}

		for _, srcPath := range possiblePaths {
			deDupFile(srcPath, cmpPath)
		}
	}
}

func deDup(srcPath, cmpPath string) {
	srcSize2FilePaths := BuildFileSizeMap(srcPath)
	cmpSize2FilePaths := BuildFileSizeMap(cmpPath)

	for size, srcPaths := range srcSize2FilePaths {
		if size == 0 {
			continue
		}

		if cmpPaths, ok := cmpSize2FilePaths[size]; ok {
			deDupFiles(srcPaths, cmpPaths)
		}
	}
}

func DeDup(srcPath, cmpPath string) {
	// 确认文件存在
	srcFI, err := os.Stat(srcPath)
	lErr.PanicErr(err)
	cmpFI, err := os.Stat(cmpPath)
	lErr.PanicErr(err)

	// 确认参数均为文件夹
	if !srcFI.IsDir() || !cmpFI.IsDir() {
		panic("Args require directory.")
	}

	// 转化为绝对路径方便操作
	// 主要是防止传入的两个参数有父子关系
	srcPath, err = filepath.Abs(srcPath)
	lErr.PanicErr(err)
	cmpPath, err = filepath.Abs(cmpPath)
	lErr.PanicErr(err)

	deDup(srcPath, cmpPath)
}
