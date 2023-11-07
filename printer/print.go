package printer

import (
	"fmt"
	"time"
)

func PrintHelloWorld()  {
	PrintBegining()
	fmt.Println("mohammadjavad dahane maro gayidi :(")
}

func PrintBegining() {
	fmt.Print(time.Now())
	fmt.Print(" => ")
}
