package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ReadBlacklist(file string) {
	file, err := os.Open("blacklist.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
}
