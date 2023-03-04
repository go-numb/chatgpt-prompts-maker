package prompts

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

const (
	FILENAME = "./_data/prompts.csv"
)

type Act struct {
	Actor  string `csv:"act"`
	Prompt string `csv:"prompt"`
}

func Load() []Act {
	abs, _ := filepath.Abs(FILENAME)
	fmt.Println(abs)
	f, _ := os.Open(abs)
	defer f.Close()

	acts := make([]Act, 0)
	if err := gocsv.Unmarshal(f, &acts); err != nil {
		log.Println(err)
		return nil
	}

	return acts
}
