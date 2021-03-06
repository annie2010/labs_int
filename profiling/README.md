<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Profiling Lab

---
## <img src="../assets/lab.png" width="auto" height="32"/> Mission

> Implement a Fibonacci number web service that produces the first n Fibonacci numbers given n as input.

  ```text
    fib(0) = 0
    fib(1) = 1
    fib(n) = fib(n-2)+fib(n-1)
  ```

+ Clone the [Labs Repo](https://github.com/gopherland/labs_int)
+ Cd profiling
+ A Fib service initial implementation has been provided for you.
+ Using profiling technics, establish the service performance profile and baseline.
  + Record initial numbers gathered via hey or apache bench
  + Is this service CPU/Mem or IO bound?
+ Leverage MicroBenchmarks to assert your profiling experimentation and improvements.
+ What do you notice and what can you improve?
+ Tune the implementation and repeat the profiling process to produce the best possible results.

### Commands

+ Ensure the tests are passing!
+ Start your web server and check for valid output
+ Check GC pace. What do you notice?
+ Check your CPU profile

```shell
# Check endpoint
http :4500/fib n==5
# Load service using hey or apachebench
go get -u github.com/rakyll/hey
hey -c 2 -n 100000 http://localhost:4500/fib?n=20
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)