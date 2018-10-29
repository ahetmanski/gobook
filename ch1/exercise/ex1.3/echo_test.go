// Benchmark tests for echo implementations comparison.
// Run with `go test -bench=. -benchmem`
package echo

import (
	"strings"
	"testing"
)

var inputParams = strings.Split("Hello мир! It's me! Benchmark test is here once again.", " ")

func BenchmarkEchoSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoSlow(inputParams)
	}
}

func BenchmarkEchoFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoFast(inputParams)
	}
}
