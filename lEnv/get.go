package lEnv

import "os"

func Get(key, _default string) string {
	res := os.Getenv(key)
	if res == "" {
		res = _default
	}

	return res
}
