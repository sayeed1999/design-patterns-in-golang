package main

import (
	"fmt"
	"time"

	"github.com/sayeed1999/design-patterns-in-golang/proxy-pattern/image"
	proxyimage "github.com/sayeed1999/design-patterns-in-golang/proxy-pattern/proxy-image"
)

func main() {

	fmt.Printf("\nWe are instantiating real image but not calling display...\n")
	time.Sleep(time.Second * 1)
	realImage := image.NewRealImage("abc.jpeg")
	time.Sleep(time.Second * 1)

	fmt.Printf("\nWe are calling display...\n")
	time.Sleep(time.Second * 1)
	realImage.Display()
	time.Sleep(time.Second * 1)

	fmt.Printf("\nWe are instantiating proxy image but not calling display...\n")
	time.Sleep(time.Second * 1)
	proxyimage := proxyimage.NewProxyImage("abc.jpeg")
	time.Sleep(time.Second * 1)

	fmt.Printf("\nWe are calling display...\n")
	time.Sleep(time.Second * 1)
	proxyimage.Display()
	time.Sleep(time.Second * 1)
}

// Benefits of this approach: -
// Without proxy, image is loaded instantly before calling Display()
// With proxy, image is not loaded until display() is called..
