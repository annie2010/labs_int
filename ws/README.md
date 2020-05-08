<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Web Service Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> Turn the book grep into a web service.

> Create a grep service that will accept a book and a word and returns a JSON response
> representing the number of occurrences of that word.

1. Clone the [Labs Repo](https://github.com/gopherland/labs_int)
1. Cd ws
1. Edit internal/handler/wc.go and implement a GrepHandler
   1. The handler gets a book and a word from the request url
   2. Use the countWords function to compute the number of occurrences
   3. Return a json response that contains the following fields:
      1. book
      2. word
      3. count
   4. Write a test (wc_test.go) to make sure your handler is working correctly
1. Edit main.go and define a gorilla mux route to your GrepHandler
   1. Use: /v1/grep/book/word as your endpoint
   2. Next define a logging middleware to log incoming requests
      1. Use the [Gorilla Handlers Repo](https://github.com/gorilla/handlers)
1. Make sure all tests are cool!
1. Launch your service and make sure your endpoint is working correctly!

## Expectations

```shell
   http :5000/v1/wc/3lpigs/pig
```

Produces...

```text
HTTP/1.1 200 OK
Content-Length: 46
Content-Type: application/json; charset=utf-8
Date: Fri, 08 May 2020 03:36:19 GMT

{
    "book": "3lpigs",
    "count": 26,
    "word": "pig"
}
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
