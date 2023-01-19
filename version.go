package _go

import (
	"fmt"
	_ "github.com/licsber/go/lEnv/dotenv/autoload"
	"os"
)

func init() {
	if os.Getenv("LICSBER") == "" {
		fmt.Println("LicsberLib " + VERSION)
	}
}

const VERSION = "v0.1.2"
