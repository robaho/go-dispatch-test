**About**

This package is a micro-benchmark which tests the performance of Go method dispatch, both interface based, and concrete type.

It has similar tests written using Java as well.

**Testing Methodology**

The tests use 'gobench' for Go and 'jmh' for Java. The tests were rung using IntelliJ support for both.

The tests were run using Go 1.11 and Java 1.8_181.

The tests were run on identical hardware and OS, iMac OS 10.13.6, core i7 3.4 ghz (4 cores, 8 threads)

Multiple runs were performed with very insignificant in the results.

**Results**


Go

```
goos: darwin
goarch: amd64
pkg: github.com/robaho/go-dispatch-test
BenchmarkIDispatch-8   	300000000	         5.75 ns/op
BenchmarkDispatch-8    	2000000000	         1.59 ns/op
PASS

```

Java

```
# JMH version: 1.21
# VM version: JDK 1.8.0_181, Java HotSpot(TM) 64-Bit Server VM, 25.181-b13
# VM invoker: /Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre/bin/java
# VM options: -Dfile.encoding=UTF-8
# Warmup: 1 iterations, 10 s each
# Measurement: 3 iterations, 1 s each
# Timeout: 10 min per iteration
# Threads: 1 thread, will synchronize iterations
# Benchmark mode: Average time, time/op
# Benchmark: test.TestJavaDispatch.TestDispatch

Benchmark                       Mode  Cnt  Score   Error  Units
TestJavaDispatch.TestDispatch   avgt    3  0.500 ± 0.120  ns/op
TestJavaDispatch.TestIDispatch  avgt    3  0.979 ± 0.097  ns/op
```

**Notes**

Interestingly, if the Go methods are declared with pointer receivers, the performance is almost 3x worse:

```
goos: darwin
goarch: amd64
pkg: github.com/robaho/go-dispatch-test
BenchmarkIDispatch-8   	100000000	        12.6 ns/op
BenchmarkDispatch-8    	2000000000	         1.57 ns/op
PASS

```