package main

import (
	"encoding/json"
	"os"

	"github.com/itchio/wharf/tlc"
)

func walk(src string) {
	info, err := tlc.Walk(src, filterDirs)
	must(err)

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(info)
}