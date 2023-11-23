# YaegiDemo

This is a demo project to show how to use [Yaegi](https://github.com/traefik/yaegi) to load and run Go code at runtime.

## How to play with this demo

Run `make run` to run the demo, which is a simple HTTP server, and then you can use `curl` to test it:

```bash
$ curl localhost:8080 -i
```

You will get the following response:

```bash
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
Via: useless-helper
Date: Thu, 23 Nov 2023 12:24:31 GMT
Content-Length: 241

Hello!
Your Content-Type is
Your User-Agent is curl/8.1.2
Your Host is localhost:8080
Your RemoteAddr is 127.0.0.1:61784
Your RequestURI is /
Your Method is GET
Your Proto is HTTP/1.1
Your URL is /
Code above is wrote by GitHub Copilot.

:)⏎
```

But, all the logic is controlled by the Go code in `rule/rule.go`, which is hot-reloaded by Yaegi at runtime.

Modify the code in `rule/rule.go` and save it, then you will see the response is changed, without restarting the server.

Have fun!
