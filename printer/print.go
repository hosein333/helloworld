package printer

import (
	"fmt"
	"time"
)

func PrintHelloWorld()  {
	PrintBegining()
	fmt.Println("Hello World")
}

func PrintBegining() {
	fmt.Print(time.Now())
	fmt.Print(" => ")
}
