**About**

This package is a micro-benchmark which tests the performance of Go method dispatch, both interface based, and concrete type.

It has similar tests written using Java as well.

**Testing Methodology**

The tests use 'gobench' for Go and 'jmh' for Java. The tests were rung using IntelliJ support for both.

The tests were run using Go 1.11 and Java 1.8_181.

The tests were run on identical hardware and OS, iMac OS 12.7.5, 4 GHz Quad-Core Intel Core i7

Multiple runs were performed with very insignificant in the results.

**Results**

Go 1.22.4
```
goos: darwin
goarch: amd64
pkg: github.com/robaho/go-dispatch-test
cpu: Intel(R) Core(TM) i7-6700K CPU @ 4.00GHz
BenchmarkIDispatch-8      	535263746	         2.118 ns/op
BenchmarkIDispatchPtr-8   	675721244	         1.773 ns/op
BenchmarkDispatch-8       	1000000000	         0.6953 ns/op
BenchmarkArrayParms-8     	  184212	      6440 ns/op
PASS

```

Java build 21+35-LTS-2513

```
Benchmark                           Mode  Cnt     Score       Error  Units
TestJavaDispatch.TestArrayDispatch  avgt    3  4775.994 ± 13310.309  ns/op
TestJavaDispatch.TestDispatch       avgt    3     0.941 ±     0.088  ns/op
TestJavaDispatch.TestIDispatch      avgt    3     1.331 ±     0.146  ns/op

```

**Notes**

Interestingly, if the Go methods are declared with pointer receivers, the performance improves - which is a change
over previous Go versions where the performance was almost 3x worse.