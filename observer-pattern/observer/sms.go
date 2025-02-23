package observer

import "fmt"

type SMS struct{} // implements Observer

func (e *SMS) Update() {
	fmt.Println("Sending SMS...")
}
