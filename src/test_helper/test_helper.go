package test_helper

import (
	"fmt"
	"strings"
	"time"
	"reflect"
)

func AssertEqual(actual interface{}, expected interface{}) {
	if actual != expected {
		panic(fmt.Sprintf("Expected %v to equal %v", actual, expected))
	}
}

func AssertNotEqual(actual interface{}, expected interface{}) {
	if actual == expected {
		panic(fmt.Sprintf("Expected %v to not equal %v", actual, expected))
	}
}

func AssertContains(container string, containee string) {
	if !strings.Contains(container, containee) {
		panic(fmt.Sprintf("Expected %v to contain  %v", container, containee))
	}
}

func AssertTrue(result bool) {
	if result != true {
		panic(fmt.Sprintf("Expected result to be true, but got %v", result))
	}
}

func AssertFalse(result bool) {
	if result != false {
		panic(fmt.Sprintf("Expected result to be false, but got %v", result))
	}
}

func AssertValueIsNil(actual interface{}){
	if !reflect.ValueOf(actual).IsNil(){
		panic(fmt.Sprintf("Expected value of %v to be nil", actual))
	}
}

func Fail(msg string) {
	panic(msg)
}

func SleepQuick() {
	time.Sleep(100 * time.Millisecond)
}