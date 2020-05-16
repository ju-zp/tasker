package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_server(t *testing.T) {
	resp, err := http.Get("http://localhost:3003/ping")

	fmt.Println(err)

	if err != nil {
		t.Fatal("Expected to get ping endpoint")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

	if err != nil {
		t.Fatal("server responded with incorrect response")
	}

	t.Log(body)
}
