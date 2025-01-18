# url_shortener
a simple url shortener service written in Go

## Stack

- web framework: `Fiber`
- DB: `PostgreSQL`
- DB toolchain: `migrate`, `sqlc`
- Container/Orchestration: Docker/Docker compose
- stress test: `wrk`

## Stress test results

### V1. Use auto-increment key

3000 - 3500 QPS

```bash
<@url_shortener>-<⎇ main>-> wrk -t2 -d30s -c10 -s ./scripts/wrk/shorten.lua http://localhost:3000/
Running 30s test @ http://localhost:3000/
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.17ms    1.17ms  16.14ms   61.95%
    Req/Sec     1.59k   584.07     3.24k    80.83%
  95024 requests in 30.01s, 12.96MB read
Requests/sec:   3166.17
Transfer/sec:    442.15KB
<@url_shortener>-<⎇ main>-<±>-> wrk -t2 -d30s -c10 -s ./scripts/wrk/shorten.lua http://localhost:3000/
Running 30s test @ http://localhost:3000/
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.81ms    1.14ms  12.12ms   66.73%
    Req/Sec     1.79k   663.18     3.16k    70.00%
  106860 requests in 30.01s, 14.57MB read
Requests/sec:   3561.01
Transfer/sec:    497.29KB
```

## Thoughts

- I chose `Fiber` and `sqlc` because in another showcase project `secumon` I have
used `Gin` and `Gorm`. It is to show that I can adapt with different tools
based on the existent project setup.
