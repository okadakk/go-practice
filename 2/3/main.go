package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	writer := gzip.NewWriter(w)

	// multiWriterを使うと、第一引数のwriterに出力されるものと同じものを、第二引数以降のio.Writerで出力できる
	multiWriter := io.MultiWriter(writer, os.Stdout)

	source := map[string]string{"Hello": "World"}

	// jsonEncoderでwriterをEncode
	err := json.NewEncoder(multiWriter).Encode(source)
	if err != nil {
		panic(err)
	}

	writer.Flush()
	writer.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
