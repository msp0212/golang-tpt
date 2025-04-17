# concepts

## spec
[https://go.dev/ref/spec]

## memory model
[https://go.dev/ref/mem]

## concurrency model
[https://go.dev/talks/2012/concurrency.slide#1]  
[https://go.dev/doc/codewalk/sharemem/]  
[https://www.youtube.com/watch?v=QDDwwePbDtw]  
[https://medium.com/@mail2rajeevshukla/unlocking-the-power-of-goroutines-understanding-gos-lightweight-concurrency-model-3775f8e696b0]  

## garbage collection
[https://tip.golang.org/doc/gc-guide]

## stateful goroutines

using `channel` and `select`

go promotes the idea of sharing memory by communicating over channels


## slices
[https://go.dev/blog/slices-intro]

---

# idioms

## errors
go does not have a concept of exceptions
Use error-indicating return values wherever possible.

---
# assignments

## worker pools using goroutines and channels

## rate limiter
 using `goroutine`, `channel` and `tickers`

 [https://gobyexample.com/rate-limiting]

---

