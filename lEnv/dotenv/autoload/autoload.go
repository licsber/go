package autoload

import (
	"github.com/licsber/go/lEnv/dotenv"
	"os"
	"path/filepath"
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		// Without overload, load local .env first.
		_ = dotenv.Load(".env", filepath.Join(home, ".env"))
	} else {
		_ = dotenv.Load()
	}
}
