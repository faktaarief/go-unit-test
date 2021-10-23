package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	// if runtime.GOOS == "windows" {
	// 	t.Skip("Can not run on Winodws")
	// }

	result := HelloWorld("Fakta")
	require.Equal(t, "Hello Fakta", result, "Result must be 'Hello Fakta'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fakta")
	// Assert using Fail function
	// assert.Equal(t, "Hello Fakta", result, "Result must be 'Hello Fakta'")

	// Require using Fail Now function
	require.Equal(t, "Hello Fakta", result, "Result must be 'Hello Fakta'")

	fmt.Println("TestHelloWorld with Assert Done")
}

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
