package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"immudb-demo/internal/api"
	model "immudb-demo/internal/model"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	_ "github.com/codenotary/immudb/pkg/stdlib"
)

type cfg struct {
	IpAddr   string
	Port     string
	Username string
	Password string
	DBName   string
}

func parseConfig() (c cfg) {
	flag.StringVar(&c.IpAddr, "addr", "immudb", "IP address of immudb server")
	flag.StringVar(&c.Port, "port", "3322", "Port number of immudb server")
	flag.StringVar(&c.Username, "user", "immudb", "Username for authenticating to immudb")
	flag.StringVar(&c.Password, "pass", "immudb", "Password for authenticating to immudb")
	flag.StringVar(&c.DBName, "db", "defaultdb", "Name of the database to use")
	flag.Parse()
	return
}

func main() {
	c := parseConfig()
	var auth = gin.Accounts{c.Username: c.Password}

	dsn := fmt.Sprintf("immudb://%s:%s@%s:%s/%s?sslmode=disable", c.Username,
		c.Password, c.IpAddr, c.Port, c.DBName)

	repo, err := model.NewMyLogImmudbRepo(context.Background(), dsn)
	if err != nil {
		os.Exit(1)
	}

	// Create myLog table if not exists
	err = repo.CreateLogTable()
	if err != nil {
		os.Exit(1)
	}

	router := gin.Default()

	// Set getRepo to context for obtaining model.MyLogRepository in handlers
	router.Use(func(context *gin.Context) {
		context.Set("dsn", dsn)
		context.Set("getRepo", func(c *gin.Context) (model.MyLogRepository, error) {
			return model.NewMyLogImmudbRepo(c, c.GetString("dsn")) //nolint:wrapcheck
		})
	})

	// Setup endpoints
	gr := router.Group("api/v1")
	gr.GET("/log", gin.BasicAuth(auth), api.FetchLog)
	gr.POST("/log", gin.BasicAuth(auth), api.InsertLog)
	gr.POST("/logs", gin.BasicAuth(auth), api.InsertLogBatch)

	// Serve static
	router.Use(static.Serve("/", static.LocalFile("./public", true)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	router.Run()
}
