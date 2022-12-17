package lolicon

import (
	"encoding/json"
	"fmt"
	"github.com/licsber/go/lNet"
)

const _v2Url = "https://api.lolicon.app/setu/v2"

var V2Config = map[string]interface{}{
	"r18":   2,  // 是否包含R18图片 0为否 1为是 2为混合
	"num":   20, // 获取结果数量
	"size":  []string{"original", "regular", "small", "thumb", "mini"},
	"proxy": "", // i.pixiv.re
}

func V2(c *lNet.Client, b []byte) map[string]interface{} {
	b, err := c.PostJSON(_v2Url, b)
	if err != nil {
		fmt.Println("Network error. " + err.Error())
		return nil
	}

	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("Parse fail. " + string(b))
		return nil
	}

	return res
}
