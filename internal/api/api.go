package api

import (
	"fmt"
	"net/http"
	"strconv"

	"immudb-demo/internal/model"

	"github.com/gin-gonic/gin"
)

func InsertLog(c *gin.Context) {
	repo, _ := c.Get("getRepo")
	r, err := repo.(func(*gin.Context) (model.MyLogRepository, error))(c)

	var log *model.Log
	if err := c.BindJSON(&log); err != nil {
		return
	}

	err = r.InsertLog(log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%s", err)})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func InsertLogBatch(c *gin.Context) {
	repo, _ := c.Get("getRepo")
	r, err := repo.(func(*gin.Context) (model.MyLogRepository, error))(c)

	var logs []*model.Log
	if err := c.BindJSON(&logs); err != nil {
		return
	}

	//TODO Could be improved - single insert query with all the data or execute in transaction
	for _, log := range logs {
		err = r.InsertLog(log)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%s", err)})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func FetchLog(c *gin.Context) {
	repo, _ := c.Get("getRepo")
	r, err := repo.(func(*gin.Context) (model.MyLogRepository, error))(c)

	lastXString := c.Query("lastx")
	var lastX int
	if len(lastXString) > 0 {
		lastX, _ = strconv.Atoi(lastXString)
	}

	res, err := r.FetchLog(lastX)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("%s", err)})
		return
	}

	c.JSON(http.StatusOK, res)
}
