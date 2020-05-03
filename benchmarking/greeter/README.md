<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Benchmark greeter Lab...

## <img src="../../assets/lab.png" width="auto" height="32"/> Mission

> Congratulations! You've just got a new gig working at Google!!
> Your PM says we have to speed up the initial greeting page load that's
> being hit by billions of users! The page greets the users with
> a user name and age.

* Using Benchmark tests implement a second version of the greeter and compare with this initial version.
* Make sure your new implementation works the same as the old one ie your tests pass for both implementations!
* HINT! Another way to concatenate strings is by using the + operator or use strings.Builder
* What are your findings?


## Up And Running...

1. Run the tests

```shell
go test
```

1. Run Benchmarks

```shell
go test --bench Greet --benchmem --benchtime=3s
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)