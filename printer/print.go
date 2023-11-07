package printer

import (
	"fmt"
	"time"
)

func PrintHelloWorld()  {
	PrintBegining()
	fmt.Println("test")
}

func PrintBegining() {
	fmt.Print(time.Now())
	fmt.Print(" => ")
}
