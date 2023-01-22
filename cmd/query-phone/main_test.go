package main

import (
	"fmt"
	"testing"
)

func init() {
	Debug()
}

func BenchmarkQuery(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		var i = 0
		for p.Next() {
			i++
			_, err := Query(fmt.Sprintf("%s%d%s", "1897", i&10000, "45"))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func TestQueryLonger(t *testing.T) {
	_, err := Query("13580198235123123213213")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestQueryShorter(t *testing.T) {
	_, err := Query("1300")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestQueryV(t *testing.T) {
	pr, err := Query("1703576")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pr)
}

func TestQueryError(t *testing.T) {
	_, err := Query("10074872323")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestQueryChar(t *testing.T) {
	_, err := Query("asd32323")
	if err == nil {
		t.Fatal("错误的结果")
	}
	t.Log(err)
}

func TestQueryCorrect(t *testing.T) {
	//1952947,广西,玉林,移动
	info, err := Query("1952947")
	if err != nil {
		t.Fatal("错误的结果")
	}
	t.Log(info)
	str := fmt.Sprintf("%s,%s,%s,%s", info.PhoneNum, info.Province, info.City, info.CardType)
	if str != "1952947,广西,玉林,中国移动" {
		t.Fatal("验证失败")
	}
}
