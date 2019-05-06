package greeting

import "testing"

// Define a function named ScrambledEggs that takes no arguments,
// and returns a string.
// In other words, define a function with the following signature:
// ScrambledEggs() string

func TestScrambledEggs(t *testing.T) {
	expected := "Scrambled eggs without bacon?!"
	if observed := ScrambledEggs(); observed != expected {
		t.Fatalf("ScrambledEggs() = %v, want %v", observed, expected)
	}
}

// BenchmarkScrambledEggs() is a benchmarking function. These functions follow the
// form `func BenchmarkXxx(*testing.B)` and can be used to test the performance
// of your implementation. They may not be present in every exercise, but when
// they are you can run them by including the `-bench` flag with the `go test`
// command, like so: `go test -v --bench . --benchmem`
//
// You will see output similar to the following:
//
// BenchmarkScrambledEggs   	2000000000	         0.46 ns/op
//
// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
//
// While benchmarking can be useful to compare different iterations of the same
// exercise, keep in mind that others will run the same benchmarks on different
// machines, with different specs, so the results from these benchmark tests may
// vary.
func BenchmarkScrambledEggs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ScrambledEggs()
	}
}
