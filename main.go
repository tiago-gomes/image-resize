package main

import (
	"fmt"
	Image "imageprocessing/imageprocessing"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: go run image.go resize <path> <width> <height>")
		return
	}

	action := args[1]
	if action != "resize" {
		fmt.Println("Invalid action, available actions are: resize")
		return
	}

	filePath := args[2]
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
		return
	}

	widthStr := args[3]
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		fmt.Println("Invalid width.")
		return
	}

	heightStr := args[4]
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Println("Invalid height.")
		return
	}

	Image.Resize(filePath, width, height)
}
