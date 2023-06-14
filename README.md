# hertz-jsonp

JSONP Middleware for Hertz Framework

### install

```
go get github.com/li-jin-gou/hertz-jsonp
```

### usage

```go
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	jsonp "github.com/li-jin-gou/hertz-jsonp"
)

func main() {
	r := server.Default()
	r.Use(jsonp.JsonPHertz())
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, utils.H{
			"message": "pong",
		})
	})
	r.Spin()
}
```

**output:**

```
curl -v http://127.0.0.1:8080/ping\?callback\=callback

*   Trying 127.0.0.1...
> GET /ping?callback=callback HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.43.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Sun, 28 Aug 2016 05:52:00 GMT
< Content-Length: 28
< 
* Connection #0 to host 127.0.0.1 left intact
callback({"message":"pong"})
```

