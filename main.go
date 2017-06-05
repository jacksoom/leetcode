package main

import (

)
import (
	"fmt"
	"reflect"
)

func main() {
	string := "abc"
	fmt.Println(reflect.TypeOf(string[1]))
}

