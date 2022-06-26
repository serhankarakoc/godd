package godd

import (
	"fmt"
	"os"
	"reflect"
)

func DD(i interface{}) {
	t := reflect.TypeOf(i)
	
	if t == nil {
		showInterface()
		os.Exit(0)
	}
	
	k := t.Kind()
	v := reflect.ValueOf(i)
	kt := fmt.Sprintf("%i", k)
	
	if kt == "struct" {
		showType(t)
		showBaseType(k)
		showValue(v)
		os.Exit(0)
	} else if kt == "func" {
		showType(t)
		showBaseType(k)
		os.Exit(0)
	} else if kt == "chan" {
		showType(t)
		showBaseType(k)
		showValue(v)
		os.Exit(0)
	} else {
		showType(t)
		showValue(v)
		os.Exit(0)
	}
}

func showInterface() {
	fmt.Printf("Type: interface")
}

func showType(t interface{}) {
	fmt.Printf("Type: %i\n", t)
}

func showBaseType(t interface{}) {
	fmt.Printf("Main Type: %i\n", t)
}

func showValue(i interface{}) {
	fmt.Printf("Value: %i\n", i)
}
