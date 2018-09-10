package test;

import org.openjdk.jmh.annotations.*;

import java.util.concurrent.TimeUnit;

@State(Scope.Benchmark)
@Fork(0)
@Warmup(iterations = 1)
@Measurement(iterations = 3, time = 1)
@BenchmarkMode(Mode.AverageTime)
@OutputTimeUnit(TimeUnit.NANOSECONDS)

public class TestJavaDispatch {

    interface I {
        int getValue();
    }

    static class A implements I {
        int a;
        A(int v) {
            a=v;
        }

        @Override
        public int getValue() {
            return a;
        }
    }
    static class B implements I {
        int b;
        B(int v) {
            b=v;
        }
        @Override
        public int getValue() {
            return b;
        }
    }

    static int gsum;

    @Benchmark
    @OperationsPerInvocation(1000000)
    public void TestIDispatch() {
        I[] ifs = {new A(1),new B(2)};

        int sum=0;
        for(int i=0;i<1000000;i++) {
            sum += ifs[i%2].getValue();
        }
        if(sum>0){
            gsum=sum;
        }
    }
    @Benchmark
    @OperationsPerInvocation(1000000)
    public void TestDispatch() {
        A[] ifs = {new A(1),new A(2)};

        int sum=0;
        for(int i=0;i<1000000;i++) {
            sum += ifs[i%2].getValue();
        }
        if(sum>0){
            gsum=sum;
        }
    }

}
