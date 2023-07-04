package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"immudb-demo/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		context.Set("getRepo", func(c *gin.Context) (model.MyLogRepository, error) {
			return MyLogRepoMock{}, nil
		})
	})
	return router
}

func TestInsertLog(t *testing.T) {
	r := setupRouter()
	r.POST("/log", InsertLog)
	log := model.Log{
		Data:   "random",
		Source: "s1",
	}
	jsonValue, _ := json.Marshal(log)
	req, _ := http.NewRequest("POST", "/log", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestInsertLogBatch(t *testing.T) {
	r := setupRouter()
	r.POST("/logs", InsertLogBatch)
	logs := []*model.Log{
		{
			Data:   "random 1",
			Source: "s1",
		},
		{
			Data:   "random 2",
			Source: "s2",
		},
	}
	jsonValue, _ := json.Marshal(logs)
	req, _ := http.NewRequest("POST", "/logs", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFetchLog(t *testing.T) {
	r := setupRouter()
	r.GET("/log", FetchLog)

	t.Run("Get logs", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/log", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Get last x logs", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/log?lastx=1", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var logs []*model.Log
		json.Unmarshal(w.Body.Bytes(), &logs)
		assert.Equal(t, 1, len(logs))
	})
}

type MyLogRepoMock struct {
}

func (m MyLogRepoMock) FetchLog(lastX int) ([]*model.Log, error) {
	res := []*model.Log{
		{Data: "Some data 1", Source: "s1"},
		{Data: "Some data 2", Source: "s2"},
		{Data: "Some data 3", Source: "s3"},
		{Data: "Some data 4", Source: "s4"},
	}
	if lastX != 0 {
		return res[len(res)-lastX:], nil
	}
	return res, nil
}
func (m MyLogRepoMock) InsertLog(log *model.Log) error {
	return nil
}

func (m MyLogRepoMock) CreateLogTable() error {
	return nil
}

func (m MyLogRepoMock) Close() error {
	return nil
}
