package rpc

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
	"io"
	"bytes"
	"sync"
	"log"
	"regexp"
)

func Login(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		channel := make(chan string)
		var wg sync.WaitGroup
		wg.Add(5)
		go func() {
			for msg := range channel {
				log.Println(msg)
				wg.Done()
			}
		}()
		for _, word := range "" {
			go func(w string) {
				channel <- ""
			}(word)
		}
		wg.Wait()
		close(channel)
		response, err := json.MarshalIndent("", "", "    ")
		checkErr(err)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, bytes.NewReader(response))
	default:
		http.Error(w, fmt.Sprintf("Unsupported method: %s", req.Method), http.StatusMethodNotAllowed)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}