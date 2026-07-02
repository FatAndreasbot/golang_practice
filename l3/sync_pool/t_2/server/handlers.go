package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var requestDataPool sync.Pool = sync.Pool{
	New: func() any {
		return &RequestData{
			data: make(map[string]any),
		}
	},
}

func NewServer() *Server {
	var server Server

	server = append(server, Handler{
		path: "/",
		handler: func(w http.ResponseWriter, r *http.Request) {
			reqestData := requestDataPool.Get().(*RequestData)
			defer reqestData.Reset()
			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(&reqestData.data)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintln(w, "error during json parsing")
			}

			for key, value := range reqestData.data {
				fmt.Fprintf(w, "%s : %v\n", key, value)
			}

		},
	})

	return &server
}
