package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rajatjindal/wasi-go-sdk/pkg/wasihttp"
)

func main() {}
func init() {
	wasihttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		client := wasihttp.NewClient()
		req, err := http.NewRequest(http.MethodGet, "https://random-data-api.fermyon.app/animals/json", nil)
		if err != nil {
			fmt.Println("failed to create request", err)
			os.Exit(1)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("failed to make outbound request", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("expected status code: %d, got: %d\n", http.StatusOK, resp.StatusCode)
			os.Exit(1)
		}

		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("failed to read response body", err)
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(raw)
	})
}
