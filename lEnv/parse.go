package lEnv

import "os"

func Parse(requires []string) map[string]string {
	res := make(map[string]string)
	for _, r := range requires {
		res[r] = os.Getenv(r)
		if res[r] == "" {
			panic("请设置环境变量: " + r)
		}
	}

	return res
}
