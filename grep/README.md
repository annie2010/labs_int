<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Testing and Benchmarking Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Mission


> Implement Grep+Wc! Your CLI application should take a word and a file name and
> report back the number of occurrences of the word in the file.

* Implement and test a `WordCount` function that computes a given word occurence count.
* Benchmark your count function.
* Using `benchstat` checkout to variation and make sure you have a solid measurement.
* Can you speed up your initial implementation?
* Implement and test a second implementation.
* Using `benchstat` compare your first and second implementations.
* BONUS! Can you think of a more performant way to implement the count? How does your solution compares to using unix grep?

### Setup

```shell
# Install benchstat
go get -u golang.org/x/perf/cmd/benchstat
# Navigate to your own Go workspace
$ cd $HOME/gopherland
# Clone Labs Repo if not done already
$ git clone https://github.com/gopherland/ntnx_labs.git
# Lab dir
cd benchmark/grep
# Get cracking!!
```

### Commands

```shell
# Run your program
go run cmd/main.go # => Found 90 occurrences of "moby"
# Run a benchmark
go test --run xxx --bench Count --count 10
# Compare implementations
go test --run xxx --bench Count --count 10 | tee v1.txt
go test --run xxx --bench Count --count 10 | tee v2.txt
benchstat v1.txt v2.txt
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)