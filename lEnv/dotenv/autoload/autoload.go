package autoload

import (
	"github.com/licsber/go/lEnv/dotenv"
	"os"
	"path/filepath"
)

func init() {
	_ = dotenv.Load()
	home, err := os.UserHomeDir()
	if err != nil {
		// Without overload, load local .env first, then home.
		_ = dotenv.Load(filepath.Join(home, ".env"))
	}
}
