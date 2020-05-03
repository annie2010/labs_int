<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Documentation Lab...

## <img src="../../assets/lab.png" width="auto" height="32"/> Mission

> FizzBuzz Calculator - Implement, test and document a Fizzbuzz calculator
> ready to be shared with the rest of the GO community. No pressure right?

* In the directory documentation/fizzbuzz
  * Change the module file to reflect your own github user
  * For example: github.com/MY_GITHUB_HANDLE/fizzbuzz
  * Make sure to export and document your compute function
* In your test make sure you use the fizzbuzz_test package as you will need to
  export your compute function.
* Add example documentation and example tests!

### Checklist

To prepare for your release make sure your **fizzbuzz** package is:

* Fully tested and a 100% covered using the table test technic
* Fully documented with examples usage!

### Reminders

```shell
# Run your tests
go test
# Checkout coverage
go test -coverprofile cov.out && go tool cover -func cov.out
# Check you're going fast!
go test -v --run xxx -bench Fizz --benchmem
# Checkout your package documentation
go doc fizzbuzz compute
```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)