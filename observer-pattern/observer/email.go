package observer

import "fmt"

type Email struct{} // implements Observer

func (e *Email) Update() {
	fmt.Println("Sending email...")
}
