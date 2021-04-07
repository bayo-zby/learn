package interview

import "fmt"

func PanicTest() string {
	fmt.Println("1")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic")

}
