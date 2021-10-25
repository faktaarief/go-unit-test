package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/** Benchmark testing
* Running All Benchmark and All Unit Test on Root
* $ go test -v -bench=. ./...
*
* Running All Benchmark Without Unit Test on Root
* $ go test -v -run=NotMatchUnitTest -bench=. ./...
*
* Running All Benchmark and All Unit Test
* $ go test -v -bench=.
*
* Only Running All Benchmark Without Unit Test
* $ go test -v -run=NotMatchUnitTest -bench=.
*
* Only Running Specific Benchmark Without Unit Test
* $ go test -v -run=NotMatchUnitTest -bench=BenchmarkHelloWorldArief
 */

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Fakta",
			request: "Fakta",
		},
		{
			name:    "Arief",
			request: "Arief",
		},
		{
			name:    "Farhanto",
			request: "Farhanto",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//  Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Fakta", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fakta")
		}
	})
	b.Run("Arief", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Arief")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fakta")
	}
}

func BenchmarkHelloWorldArief(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Arief")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fakta",
			request:  "Fakta",
			expected: "Hello Fakta",
		},
		{
			name:     "Arief",
			request:  "Arief",
			expected: "Hello Arief",
		},
		{
			name:     "Farhanto",
			request:  "Farhanto",
			expected: "Hello Farhanto",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	/*
	* How to Running Test only on Sub Test (Benchmark Same Syntax)
	* $ go test -v -run=TestSubTest/Fakta
	*
	* How to Running Test contains "Fakta"
	* $ go test -v -run=/Fakta
	**/
	t.Run("Fakta", func(t *testing.T) {
		result := HelloWorld("Fakta")
		require.Equal(t, "Hello Fakta", result, "Result must be 'Hello Fakta'")
	})
	t.Run("Arief", func(t *testing.T) {
		result := HelloWorld("Arief")
		require.Equal(t, "Hello Arief", result, "Result must be 'Hello Arief'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

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
