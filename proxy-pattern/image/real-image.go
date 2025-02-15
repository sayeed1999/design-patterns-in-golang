package image

import "fmt"

// RealImage implements Image
type RealImage struct {
	filename string
}

// NewRealImage is the constructor for realImage
func NewRealImage(filename string) *RealImage {
	img := &RealImage{filename: filename}
	img.loadImageFromDisk() // Load when created
	return img
}

func (r *RealImage) loadImageFromDisk() {
	fmt.Printf("Loading image from disk: %s...\n", r.filename)
	fmt.Println("...")
}

func (r *RealImage) Display() {
	fmt.Println("Displaying image from realImage: " + r.filename)
}
