<img src="../../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>


# Deployment Lab

---
## <img src="../../assets/lab.png" width="auto" height="32"/> Your Mission...

> In this lab, we are going to fetch books and count words concurrently.
> Each worker will take a book and a word to report on. Once
> completed the worker will update a map of books and report its word count.
> Workers will need to write to a shared book counter location 😜

1. In your cloned labs repo, navigate to concurrency/mutex directory.
2. Implement a worker that takes a book url and a word and computes the word count for that book.
3. If a worker successfully got a book from the external repo and counted its word, it updates the
   global book counter.
4. If a worker fails, it must make sure to zero out the book counter for that book and word
5. When completed, print out a summary of all books scanned including the word list counts for each of
   the books.
6. Make sure no data races exist in your program!


## Commands

   ```shell
   # Check for race conditions...
   go run --race main.go
   ```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)