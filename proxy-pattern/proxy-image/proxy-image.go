package proxyimage

import (
	"fmt"

	"github.com/sayeed1999/design-patterns-in-golang/proxy-pattern/image"
)

// proxyImage implements Image
type proxyImage struct {
	realImage *image.RealImage // Use concrete type here
	filename  string
}

// NewProxyImage is the constructor for proxyImage
func NewProxyImage(filename string) *proxyImage {
	return &proxyImage{filename: filename}
}

func (p *proxyImage) Display() {
	if p.realImage == nil {
		fmt.Println("Creating real image inside proxy...")
		p.realImage = image.NewRealImage(p.filename)
	}
	fmt.Println("Displaying image from proxy: " + p.filename)
}
