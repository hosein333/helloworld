package printer

import (
	"fmt"
	"time"
)

// PrintHelloWorld print my hello message to world
func PrintHelloWorld()  {
	PrintBegining()
	fmt.Println("mohammadjavad dahane maro gayidi :(")
}

func PrintBegining() {
	fmt.Print(time.Now())
	fmt.Print(" => ")
}
