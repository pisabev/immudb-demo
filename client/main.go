package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	host := "localhost:8080"
	err := insertLogs(host, []*Log{
		{Data: "test 1", Source: "test 1"},
		{Data: "test 2", Source: "test 2"},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := getLogs(host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, r := range res {
		fmt.Println(r.Id, r.Data, r.DateTime, r.Source)
	}
}

type Log struct {
	Id       int       `json:"id"`
	DateTime time.Time `json:"datetime"`
	Data     string    `json:"data"`
	Source   string    `json:"source"`
}

func getLogs(host string) (result []*Log, err error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("http://%v/api/v1/log", host), nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return result, err
	}
	resBody, err := io.ReadAll(res.Body)
	err = json.Unmarshal(resBody, &result)

	return
}

func insertLogs(host string, data []*Log) error {
	payload, err := json.Marshal(data)
	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%v/api/v1/logs", host), bytes.NewBuffer(payload))

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}
