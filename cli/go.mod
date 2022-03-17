module github.com/cryptotweet.io/cli

go 1.17

replace github.com/cryptotweet.io/internal/common => ../internal/common

replace github.com/cryptotweet.io/internal/tweet => ../internal/tweet

require (
	github.com/buger/jsonparser v1.1.1
	github.com/chromedp/cdproto v0.0.0-20220217222649-d8c14a5c6edf
	github.com/chromedp/chromedp v0.7.8
	github.com/cryptotweet.io/internal/common v0.0.0-00010101000000-000000000000
	github.com/cryptotweet.io/internal/tweet v0.0.0-00010101000000-000000000000
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/urfave/cli/v2 v2.3.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220308174144-ae0e22291548 // indirect
	google.golang.org/grpc v1.45.0 // indirect
)
