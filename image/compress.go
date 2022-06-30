package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func main() {
	buffer, err := bimg.Read("../assets/litter-hermit-crab/textures/Crab_BaseColor.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	newImg := bimg.NewImage(buffer)
	newBuf, err := newImg.Process(bimg.Options{
		Quality:     50,
		Compression: 1,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	bimg.Write("/", newBuf)
}
