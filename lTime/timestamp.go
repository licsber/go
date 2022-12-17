package lTime

import (
	"github.com/licsber/go/lErr"
	"time"
)

// LocalTimeStr returns 2022-10-24 11:22:33
func LocalTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ShanghaiTimeStr(t time.Time) string {
	shh, err := time.LoadLocation("Asia/Shanghai")
	lErr.PanicErr(err)

	return t.In(shh).Format("2006-01-02 15:04:05")
}
