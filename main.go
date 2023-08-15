package main

import (
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

func main() {
	http.ListenAndServe(":8080", &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct",
			"GOSUMDB=off",
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
	})
}
