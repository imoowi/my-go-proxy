package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

type discardWriter struct{}

func (discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
func main() {
	tmpDir, err := os.MkdirTemp("cache", "goproxy.cache")
	if err != nil {
		return
	}
	dirCacher := goproxy.DirCacher(tmpDir)
	http.ListenAndServe(":8080", &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct",
			"GOSUMDB=off",
		),
		Cacher:      dirCacher,
		TempDir:     tmpDir,
		ErrorLogger: log.New(&discardWriter{}, "", 0),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
	})
}
