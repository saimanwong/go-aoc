package functions

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) []string {
	var ret []string
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}
