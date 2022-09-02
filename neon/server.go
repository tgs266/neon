package neon

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/controllers"
	"github.com/tgs266/neon/neon/store"
	"github.com/tgs266/neon/ui"
)

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// check if indexing is allowed
	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string, index bool) static.ServeFileSystem {
	subFS, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(subFS),
		indexes:    index,
	}
}

func Start(host, username, password, port string, useUi bool) {
	store.CreateStore(host, username, password)
	r := gin.Default()
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"postgres": store.IsConnected(),
		})
	})
	controllers.Routes(r)
	if useUi {
		web := EmbedFolder(ui.Embedded, "dist", true)
		staticServer := static.Serve("/", web)
		r.Use(staticServer)
	}
	r.Run(":" + port)
}
