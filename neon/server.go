package neon

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tgs266/neon/neon/controllers"
	"github.com/tgs266/neon/neon/kubernetes"
	"github.com/tgs266/neon/neon/services"
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

func Start(host, username, password, port string, useUi bool, reset bool, inCluster bool, kubePath string) {
	store.CreateStore(host, username, password, reset)
	kubernetes.InitKubernetes(inCluster, kubePath)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"postgres": store.IsConnected(),
		})
	})
	if useUi {
		web := EmbedFolder(ui.Embedded, "dist", true)
		staticServer := static.Serve("/", web)
		r.Use(staticServer)
		r.NoRoute(func(c *gin.Context) {
			if c.Request.Method == http.MethodGet &&
				!strings.ContainsRune(c.Request.URL.Path, '.') &&
				!strings.HasPrefix(c.Request.URL.Path, "/api/") {
				c.Request.URL.Path = "/"
				staticServer(c)
			}
		})
	}
	controllers.Routes(r)
	services.InitPool(2)
	r.Run(":" + port)
}
