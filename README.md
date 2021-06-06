# golang-testcontainers-spike
Spike to explore using testcontainers with Go

## To run

`$ go test`

`$ go test -v 2>&1 | go-junit-report > junit-results.xml` will output test results in JUnit format

`$ go test -v -bench . -count 5 2>&1 | go-junit-report > junit-benchmark-results.xml` will output benchmark results in JUnit format
