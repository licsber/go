package lFile

import "os"

func Save(p string, b []byte) error {
	// 2022-12-12 给多点权限
	return os.WriteFile(p, b, 0777)
}
