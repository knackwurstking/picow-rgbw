package frontend

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	//go:embed dist
	dist embed.FS
)

// Dist
func Dist() http.FileSystem {
	fsys, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
