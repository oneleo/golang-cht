
# 參考
- https://morioh.com/p/bc1661ae78e8?fbclid=IwAR35GcHS7adsts9aWrfGxMz57974LQCRrkFW8e9yC2Ia--794u0SvJ6bf4U

- https://blog.wu-boy.com/2019/06/how-to-release-the-v2-or-higher-version-in-go-module/

# 安裝
- github
- go

# 要在 GOPATH 內建 hello-local 資料夾

# Init go module (if not initialized):
```go
go mod init
```

# Get package:
```go
go get -u -v github.com/maxence-charriere/go-app/v6
```

# 建立 .\app\main.go
```go
package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type hello struct {
    app.Compo
    name string
}

func (h *hello) Render() app.UI {
    return app.Div().Body(
        app.Main().Body(
            app.H1().Body(
                app.Text("Hello, "),
                app.If(h.name != "",
                    app.Text(h.name),
                ).Else(
                    app.Text("World"),
                ),
            ),
            app.Input().
                Value(h.name).
                Placeholder("What is your name?").
                AutoFocus(true).
                OnChange(h.OnInputChange),
        ),
    )
}

func (h *hello) OnInputChange(src app.Value, e app.Event) {
    h.name = src.Get("value").String()
    h.Update()
}

func main() {
    app.Route("/", &hello{})
    app.Route("/hello", &hello{})
    app.Run()
}
```

# 在 PS 編譯
```powershell
> $env:GOOS="js"
> $env:GOARCH="wasm"
> go build -o app.wasm .\app
```

# 在 CMD 編譯
```cmd
> SET GOOS=js
> SET GOARCH=wasm
> go build -o app.wasm .\app
```

# 建立 .\html\main.go
```go
package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:  "Hello Demo",
		Author: "Maxence Charriere",
	}

	if err := http.ListenAndServe(":7777", h); err != nil {
		panic(err)
	}
}
```

# 在 PS 編譯
```ps
> $env:GOOS="windows"
> $env:GOARCH="amd64"
> go build .\html
```

# 連線
- 127.0.0.1:7777