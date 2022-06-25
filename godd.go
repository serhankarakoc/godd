package godd

import (
	"fmt"
	"os"
	"reflect"
)

func DD(v interface{}) {
	t := reflect.TypeOf(v)
	
	if t == nil {
		showInterface()
		os.Exit(0)
	}
	
	k := t.Kind()
	val := reflect.ValueOf(v)
	kf := fmt.Sprintf("%v", k)
	
	if kf == "struct" {
		showType(t)
		showBaseType(k)
		showValue(val)
		os.Exit(0)
	} else if kf == "func" {
		showType(t)
		showBaseType(k)
		os.Exit(0)
	} else if kf == "chan" {
		showType(t)
		showBaseType(k)
		showValue(val)
		os.Exit(0)
	} else {
		showType(t)
		showValue(val)
		os.Exit(0)
	}
}

func showInterface() {
	fmt.Printf("Type: interface")
}

func showType(t interface{}) {
	fmt.Printf("Type: %v\n", t)
}

func showBaseType(t interface{}) {
	fmt.Printf("Underlying Type: %v\n", t)
}

func showValue(v interface{}) {
	fmt.Printf("Value: %v\n", v)
}
