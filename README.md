A server that returns the status code (or timeout) specified.

Usage:
- `go build`
- `./sampleresponse`
- `curl -H "Host: 503.status.http.example.com" http://localhost:8080`
- `time curl -H "Host: 3.timeout.http.example.com" http://localhost:8080` — should wait for 3+ seconds
- optionally, set env variable `BASE_DOMAIN` to something other than "example.com"
- optionally, set env variable `HTTP_PORT` to something other than "8080"

MIT license

mailto: ximaera@gmail.com
