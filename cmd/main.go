package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"immudb-demo/internal/api"
	model "immudb-demo/internal/model"

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
	gr.GET("/log", api.FetchLog)
	gr.POST("/log", api.InsertLog)
	gr.POST("/logs", api.InsertLogBatch)

	router.Run()
}
