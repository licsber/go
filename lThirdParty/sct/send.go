package sct

import (
	"errors"
	"fmt"
	"github.com/licsber/go/lNet"
	"net/url"
	"strings"
)

func SendPlainText(sctKey, title, text string) error {
	r := []rune(title)
	if len(r) > 32 {
		fmt.Println("标题大于32字符 将被截断.")
		text = title + "\n" + text
		title = string(r[:32])
	}

	u := fmt.Sprintf("https://sctapi.ftqq.com/%s.send", sctKey)
	data := url.Values{
		"title": {title},
		"desp":  {strings.Replace(text, "\n", "\n\n", -1)},
	}

	c := &lNet.Client{}
	res, err := c.JSON(c.PostFormBytes(u, data))
	if err != nil {
		return err
	}

	if res["code"].(float64) != 0 {
		fmt.Println(res)
		return errors.New(res["message"].(string))
	}

	return nil
}
