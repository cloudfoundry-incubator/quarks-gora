package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	. "code.cloudfoundry.org/quarks-gora"
)

func Test(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Gora))
	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	os.Setenv("FOO", "BAR")
	tests := []struct {
		name     string
		r        *http.Request
		contains string
		status   int
	}{
		{name: "1: testing get", r: newreq("GET", ts.URL+"/", nil), contains: "FOO=BAR", status: 200},
		{name: "2: testing post", r: newreq("POST", ts.URL+"/", strings.NewReader("exit 0")), contains: "OK", status: 200},
		{name: "2: testing post (error)", r: newreq("POST", ts.URL+"/", strings.NewReader("exit 1")), contains: "failed executing", status: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Fatal("Wrong status code returned")
			}

			body := resp.Body
			data, err := ioutil.ReadAll(body)
			if err != nil {
				t.Fatal(err)
			}
			result := string(data)
			if !strings.Contains(result, tt.contains) {
				t.Fatal(fmt.Sprintf("Response didn't contained: %s", tt.contains))
			}
		})
	}
}
