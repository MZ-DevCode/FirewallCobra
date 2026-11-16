package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ReadBlacklist(filename string) map[uint32]bool {
	blacklist := make(map[uint32]bool)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return blacklist
	}

	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
}
