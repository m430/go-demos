package main

import (
	"strings"

	"github.com/samber/lo"
)

func getFileList(path string) {
}

func imageCompress() {

}

func isPictureFormat(path string) (string, string) {
	temp := strings.Split(path, ".")
	if len(temp) <= 1 {
		return "", ""
	}
	exts := []string{"jpg", "png", "jpeg"}
	if lo.Count(exts, temp[1]) == 0 {
		return "", ""
	}
	return temp[0], temp[1]
}

func main() {

}
