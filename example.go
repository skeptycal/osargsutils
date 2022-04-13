// Copyright (c) 2021 Michael Treanor
// https://github.com/skeptycal
// MIT License

package osargs

import (
	"fmt"
	"log"
)

func Example() {
	here, me, err := HereMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Example import acknowledgement from package osargs")
	fmt.Println("The program name is: ", me)
	fmt.Println("The program path is: ", here)
}
