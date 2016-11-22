### Concurrent REST calls

A simple program to test concurrent REST calls in go.

For more information on Go concurrency reference 

* https://www.golang-book.com/books/intro/10 
* http://blog.narenarya.in/concurrent-http-in-go.html
* http://stackoverflow.com/questions/33104192/how-to-run-10000-goroutines-in-parallel-where-each-routine-calls-an-api
* https://medium.com/golangspec/synchronized-goroutines-part-i-4fbcdd64a4ec#.qav8o43uj
* https://medium.com/golangspec/synchronized-goroutines-part-ii-b1130c815c9d#.mximrcqzv
* http://nomad.so/2016/01/interesting-ways-of-using-go-channels/



#### Example input
```
Douglas, Fils, oceanleadership.org
Adam, Shepherd, whoi.edu
Robert, Arko, columbia.edu
John, Smith, edu
```

#### Example Output
```
Fils:ConcurrentREST dfils$ go run main.go -token=xxx-xxx-xxx-xxx
Results     John     Smith        http://orcid.org/0000-0003-1012-3584      0000-0003-1012-3584     orcid.org
Results     John     Smith        http://orcid.org/0000-0001-6066-793X      0000-0001-6066-793X     orcid.org
Results     John     Smith        http://orcid.org/0000-0001-9469-2164      0000-0001-9469-2164     orcid.org
Results     John     Smith        http://orcid.org/0000-0002-1963-4092      0000-0002-1963-4092     orcid.org
Results     John     Smith        http://orcid.org/0000-0001-5200-8811      0000-0001-5200-8811     orcid.org
Results     John     Smith        http://orcid.org/0000-0002-5028-6874      0000-0002-5028-6874     orcid.org
Results     John     Smith        http://orcid.org/0000-0003-3474-6292      0000-0003-3474-6292     orcid.org
Results     John     Smith        http://orcid.org/0000-0003-0643-3918      0000-0003-0643-3918     orcid.org
Results     John     Smith        http://orcid.org/0000-0001-9684-8847      0000-0001-9684-8847     orcid.org
Results     John     Smith        http://orcid.org/0000-0002-2451-778X      0000-0002-2451-778X     orcid.org
Results     Adam     Shepherd        http://orcid.org/0000-0003-4486-9448      0000-0003-4486-9448     orcid.org
Results     Robert     Arko        http://orcid.org/0000-0002-8278-3998      0000-0002-8278-3998     orcid.org
Results     Douglas     Fils        http://orcid.org/0000-0002-2257-9127      0000-0002-2257-9127     orcid.org
Results     Jamie     Allen        http://orcid.org/0000-0001-9550-1200      0000-0001-9550-1200     orcid.org
0.90s elapsed

```