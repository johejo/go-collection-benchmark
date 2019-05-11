# go collection benchmark

## Run

```bash
make test
```

## Result

```
go test . -v
=== RUN   TestIntersect
129
slice intersect: 251.173115ms
--- PASS: TestIntersect (0.25s)
=== RUN   TestSetIntersect
146
set intersect: 3.387332ms
--- PASS: TestSetIntersect (0.00s)
PASS
ok      github.com/johejo/go-collection-benchmark
```