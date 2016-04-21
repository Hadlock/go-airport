package status_api

import (
	"os"
	"net"
	"net/http"
	"log"
	"fmt"
	"bufio"
)

//set dockerSockPath from environment
var dockerSockPath string = os.Getenv("DOCKER_SOCKET")

func dockerContainerListHandler (w http.ResponseWriter, r *http.Request) {
	log.Printf("dockerContainerListHandler")
    // if http request is not GET
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
    
    //write header content
	w.Header().Set("Content-Type", "application/json")
    
    //if dockerSockPath is empty set to this default
	if dockerSockPath == "" {
		dockerSockPath = "/run/docker.sock"
	}
    
    
	conn, err := net.Dial("unix", dockerSockPath)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Fprintf(conn, "GET /containers/json HTTP/1.0\r\n\r\n")
		responseScanner := bufio.NewScanner(conn)
		for responseScanner.Scan() {
			if responseScanner.Text() == "" {
				responseScanner.Scan()
				fmt.Fprint(w, responseScanner.Text())
				break
			}
		}
	}
}

func Handlers () *http.ServeMux {
	log.Printf("Handlers")
	r := http.NewServeMux()
	r.HandleFunc("/docker/containers", dockerContainerListHandler)
	r.Handle("/", http.FileServer(http.Dir("static")))

	return r
}
