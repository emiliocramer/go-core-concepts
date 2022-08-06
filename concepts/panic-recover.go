package concepts

import "fmt"

// To panic is to create an error that terminates the rest of the file.
// We can call it with panic("error message")
// If we put a recover inside of a defer function however, even though the rest of the file doesn't run, the defer does

func PanicRecover() {
	fmt.Println("Everything before the panic will execute normally")

	defer func() {
		if r := recover(); r != nil {
			fmt.Print("this will print. ", r, " is stored in r")
		}
	}()

	panic("This is a panic")

	fmt.Println("This will not be printed")
}
