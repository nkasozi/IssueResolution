package http

import (
	"IssueResolution/api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Endpoints api.Endpoints
}

func (server Server) Run(port int32) {
	fmt.Printf("Starting Server")

	router, err := server.setupServer()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Server listening on port: %v", port)

	err = http.ListenAndServe(fmt.Sprintf(":%v", port), router)

	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

func (server Server) setupServer() (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/persons", server.createPerson).Methods(http.MethodPost)
	return router, nil
}

func (server Server) HttpError(errMsg string, statusCode int, writer http.ResponseWriter) {
	log.Fatalln(errMsg)
	return
}
