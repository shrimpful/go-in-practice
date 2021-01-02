package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)
var opts struct{
	Name string `short:"n" long:"name" default:"World"`
	Spanish bool `short:"s" long:"spanish"`
}

func main(){
	flags.Parse(&opts)

	if opts.Spanish==true{
		fmt.Printf("Hola %s!\n", opts.Name)
	}else{
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}