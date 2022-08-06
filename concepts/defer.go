package concepts

import "fmt"

// We can call defer to run something as the very last operation
// Usually done for readability - 1. open a file 2. defer closing it 3. do some operation on it

func Deferring() {
	defer func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
}
