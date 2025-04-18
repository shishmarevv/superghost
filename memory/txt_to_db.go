package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"superghost/pkg"
)

func main() {
	db := superghost.Open()
	superghost.CheckDB(db)
	defer superghost.Shut(db)

	root := superghost.FindProjectRoot()
	Path := filepath.Join(root, "memory", "words.txt")

	file, err := os.Open(Path)
	superghost.CheckErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		superghost.CheckErr(err)

		line = strings.TrimSpace(line)
		if len(line) > 3 {
			superghost.AddWord(db, line)
		}
	}
}
