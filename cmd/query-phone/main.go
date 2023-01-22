package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/licsber/go/lErr"
	"os"
	"path/filepath"
	"runtime"
)

var (
	content     []byte
	CardTypeMap = map[byte]string{
		CMCC:  "中国移动",
		CUCC:  "中国联通",
		CTCC:  "中国电信",
		CTCCV: "中国电信虚拟运营商",
		CUCCV: "中国联通虚拟运营商",
		CMCCV: "中国移动虚拟运营商",
	}
	totalLen, firstOffset int32
)

func version() string {
	return string(content[0:INT_LEN])
}

func totalRecord() int32 {
	return (int32(len(content)) - firstRecordOffset()) / PHONE_INDEX_LENGTH
}

func firstRecordOffset() int32 {
	return get4(content[INT_LEN : INT_LEN*2])
}

func Debug() {
	fmt.Println("Version: " + version())
	fmt.Print("TotalRecord: ")
	fmt.Println(totalRecord())
	fmt.Print("FirstRecordOffset: ")
	fmt.Println(firstRecordOffset())
}

func Query(phoneNum string) (pr *Phone, err error) {
	// 二分法查询phone数据
	if len(phoneNum) < 7 {
		return nil, errors.New("illegal phone length")
	}

	if len(phoneNum) > 11 {
		phoneNum = phoneNum[:11]
	}

	var left int32
	phoneSevenInt, err := getN(phoneNum[0 : HEAD_LENGTH-1])
	if err != nil {
		return nil, errors.New("illegal phone number")
	}

	phoneSevenInt32 := int32(phoneSevenInt)
	right := (totalLen - firstOffset) / PHONE_INDEX_LENGTH
	for {
		if left > right {
			break
		}

		mid := (left + right) / 2
		offset := firstOffset + mid*PHONE_INDEX_LENGTH
		if offset >= totalLen {
			break
		}

		curPhone := get4(content[offset : offset+INT_LEN])
		recordOffset := get4(content[offset+INT_LEN : offset+INT_LEN*2])
		cardType := content[offset+INT_LEN*2 : offset+INT_LEN*2+CHAR_LEN][0]
		switch {
		case curPhone > phoneSevenInt32:
			right = mid - 1
		case curPhone < phoneSevenInt32:
			left = mid + 1
		default:
			cByte := content[recordOffset:]
			endOffset := int32(bytes.Index(cByte, []byte("\000")))
			data := bytes.Split(cByte[:endOffset], []byte("|"))
			cardStr, ok := CardTypeMap[cardType]
			if !ok {
				cardStr = "未知电信运营商"
			}
			pr = &Phone{
				PhoneNum: phoneNum,
				Province: string(data[0]),
				City:     string(data[1]),
				ZipCode:  string(data[2]),
				AreaZone: string(data[3]),
				CardType: cardStr,
			}
			return
		}
	}

	return nil, errors.New("phone's data not found")
}

func init() {
	dir := os.Getenv("L_PHONE_DAT_DIR")
	if dir == "" {
		_, fullFilename, _, _ := runtime.Caller(0)
		dir = filepath.Dir(fullFilename)
	}

	var err error
	content, err = os.ReadFile(filepath.Join(dir, PHONE_DAT))
	lErr.PanicErr(err)

	totalLen = int32(len(content))
	firstOffset = get4(content[INT_LEN : INT_LEN*2])
}

func main() {
	argNum := len(os.Args)
	if argNum == 1 {
		// 手机号,运营商,省,市,区号,邮政编码
		fmt.Println("PhoneNumber,Carrier,Province,City,AreaZone,ZipCode,Error")
		os.Exit(0)
	}

	phoneNumber := os.Args[1]
	if argNum == 4 {
		phoneNumber += os.Args[2] + os.Args[3]
	}

	info, err := Query(phoneNumber)
	if err != nil {
		fmt.Printf("%s,err,err,err,err,err,", phoneNumber)
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("%s,%s,%s,%s,=\"%s\",=\"%s\",suc", phoneNumber, info.CardType, info.Province, info.City, info.AreaZone, info.ZipCode)
	fmt.Println()
}
