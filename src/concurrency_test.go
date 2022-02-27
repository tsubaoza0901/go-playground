package src

import "testing"

/*
goos: linux
goarch: amd64
pkg: go-playground/m/v1/src
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
BenchmarkNoneConcurrency1-2   	 2674278	       423.6 ns/op	     120 B/op	       4 allocs/op
BenchmarkNoneConcurrency2-2   	11894318	        98.74 ns/op	      48 B/op	       1 allocs/op
BenchmarkConcurrency1-2       	   86980	     13332 ns/op	     216 B/op	       5 allocs/op
BenchmarkConcurrency2-2       	  168535	      6744 ns/op	     168 B/op	       7 allocs/op
BenchmarkConcurrency3-2       	  136682	      9182 ns/op	     216 B/op	       5 allocs/op
PASS
ok  	go-playground/m/v1/src	6.805s
*/

func BenchmarkNoneConcurrency1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoneConcurrency1()
	}
}

func BenchmarkNoneConcurrency2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoneConcurrency2()
	}
}

func BenchmarkConcurrency1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concurrency1()
	}
}

func BenchmarkConcurrency2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concurrency2()
	}
}

func BenchmarkConcurrency3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Concurrency3()
	}
}
