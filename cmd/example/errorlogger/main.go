package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/skeptycal/gorepos/osargs"
)

func main() {
	fmt.Printf("%25.25s %s\n", "raw os.Args[0]:", os.Args[0])

	arg0, err := osargs.Arg0()
	if err != nil {
		log.Fatal(err)
	}

	arg0 = filepath.Base(arg0)
	fmt.Printf("%25.25s %s\n", "using Arg0():", arg0)

	here, me, err := osargs.HereMe()
	if err != nil {
		log.Fatal(err)
	}

	here = filepath.Base(here)

	fmt.Printf("%25.25s - Here: %s Me: %s\n", "using HereMe():", here, me)

}
