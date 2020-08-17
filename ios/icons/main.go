package main

import (
	"fmt"
	"strings"
)

func main() {
	icons := `512x512 iTunesArtwork
1024x1024 iTunesArtwork@2x
120x120 Icon-60@2x.png
180x180 Icon-60@3x.png
76x76 Icon-76.png
152x152 Icon-76@2x.png
167x167 Icon-83.5@2x.png
40x40 Icon-Small-40.png
80x80 Icon-Small-40@2x.png
120x120 Icon-Small-40@3x.png
29x29 Icon-Small.png
58x58 Icon-Small@2x.png
87x87 Icon-Small@3x.png`

	iconList := strings.Split(icons, "\n")
	for _, i := range iconList {
		sizeName := strings.Split(i, " ")
		size := sizeName[0]
		name := sizeName[1]
		fmt.Println(name, size)
	}
}
