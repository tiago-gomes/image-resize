package imageprocessing

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"golang.org/x/image/draw"
)

// OpenFile opens the file and returns the file pointer
func OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error opening input file: %v", err)
	}
	// Do not close the file here, return it for further use
	return file, nil
}

// Decode decodes the image from the file
func Decode(file *os.File) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Error decoding image: %v", err)
	}
	return img, nil
}

// Encode encodes the image in the specified format and writes to the outPath
func Encode(img image.Image, format, outPath string) error {
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case ".jpeg", ".jpg":
		return jpeg.Encode(file, img, nil)
	case ".png":
		return png.Encode(file, img)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}
}

// Resize resizes the image to the specified width and height
func Resize(filepath string, width int, height int) {
	if !IsAllowedExt(filepath) {
		fmt.Println("Error: File extension not allowed.")
		return
	}

	ext, err := getExtension(filepath)
	if err != nil {
		fmt.Println("Error: Invalid extension")
		return
	}

	file, err := OpenFile(filepath)
	if err != nil {
		fmt.Println("Error: Opening file failed")
		return
	}
	defer file.Close() // Close the file after it's used

	img, err := Decode(file)
	if err != nil {
		fmt.Println("Error: Decoding image failed")
		return
	}

	// Resize the image
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)

	// out path
	outPath := "output" + ext

	// Encode and save the resized image
	err = Encode(newImg, ext, outPath)
	if err != nil {
		fmt.Println("Error: Saving encoded image failed")
		return
	}

	fmt.Println("Image has been resized successfully.")
}

// getExtension returns the extension of the file
func getExtension(filePath string) (string, error) {
	ext := path.Ext(filePath)
	if ext == "" {
		return "", fmt.Errorf("Invalid extension")
	}
	return ext, nil
}

// IsAllowedExt checks if the file extension is allowed
func IsAllowedExt(filePath string) bool {
	ext := path.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}
