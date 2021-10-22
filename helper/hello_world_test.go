package helper

import (
	"fmt"
	"testing"
)

/*
* HOW TO RUN UNIT TEST
*
* Run All Unit Test
* $ go test -v ./...
*
* Run All Unit Test Inside Directory
* $ go test
*
* Run All Unit Test w/ Detail
* $ go test -v
*
* Run Specific Test
* $ go test -v -run=TestHelloWorld
*
 */

func TestHelloWorldFakta(t *testing.T) {
	result := HelloWorld("Fakta")

	if result != "Hello Fakta" {
		// error and next
		// t.Fail()

		// error and fail (RECOMMENDED)
		t.Error("Result must be Hello Fakta")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldArief(t *testing.T) {
	result := HelloWorld("Arief")

	if result != "Hello Arief" {
		// error and stop
		// t.FailNow()

		// (RECOMMENDED)
		t.Fatal("Result must be Hello Fakta")
	}

	fmt.Println("TestHelloWorld Done")
}
