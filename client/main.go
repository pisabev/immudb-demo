package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fHost := flag.String("address", "immudb:immudb@localhost:8080", "API Address/port")
	fInsert := flag.Bool("insert", false, "Inserts a random log line")
	fQueryAll := flag.Bool("query-all", false, "Fetch all logs")
	fQueryLast := flag.Int("query-last", 0, "Fetch last X logs")
	fCount := flag.Bool("count", false, "Show logs count")
	flag.Parse()

	if *fInsert {
		l := Log{Data: fmt.Sprintf("test %v", randSeq(5)), Source: fmt.Sprintf("test %v", randSeq(5))}
		err := insertLogs(*fHost, []*Log{&l})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Inserted", l)
	}

	if *fQueryAll {
		res, err := getLogs(*fHost, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, r := range res {
			fmt.Println(r.Id, r.Data, r.DateTime, r.Source)
		}
	}

	if *fQueryLast != 0 {
		res, err := getLogs(*fHost, *fQueryLast)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, r := range res {
			fmt.Println(r.Id, r.Data, r.DateTime, r.Source)
		}
	}

	if *fCount {
		res, err := getLogs(*fHost, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(len(res), "logs")
	}

}

type Log struct {
	Id       int       `json:"id"`
	DateTime time.Time `json:"datetime"`
	Data     string    `json:"data"`
	Source   string    `json:"source"`
}

func getLogs(host string, lastx int) (result []*Log, err error) {
	url := fmt.Sprintf("http://%v/api/v1/log", host)
	if lastx != 0 {
		url = fmt.Sprintf("http://%v/api/v1/log?lastx=%v", host, lastx)
	}
	req, _ := http.NewRequest("GET", url, nil)

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
