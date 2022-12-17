package autoload

import "github.com/licsber/go/lEnv/dotenv"

func init() {
	_ = dotenv.Load()
}
