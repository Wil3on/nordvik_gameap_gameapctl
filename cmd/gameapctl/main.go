package main

import (
	"os"
	"path/filepath"

	"github.com/Wil3on/nordvik_gameap_gameapctl/internal/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	os.Args[0] = filepath.Base(os.Args[0])

	app.Run(os.Args)
}
