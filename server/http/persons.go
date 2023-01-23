package http

import (
	"IssueResolution/api"
	"encoding/json"
	"net/http"
)

func (server Server) createPerson(writer http.ResponseWriter, request *http.Request) {

	var apiRequest api.CreatePersonRequest
	err := json.NewDecoder(request.Body).Decode(&apiRequest)

	if err != nil {
		server.HttpError(MsgUnmarshallRequestError, http.StatusBadRequest, writer)
		return
	}

	err, response := server.Endpoints.CreatePerson(apiRequest)

	if err != nil {
		server.HttpError(MsgInternalServerError, http.StatusInternalServerError, writer)
		return
	}

	writer.Header().Set(ContentTypeHeader, JsonContentType)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(response)

	if err != nil {
		server.HttpError(MsgInternalServerError, http.StatusInternalServerError, writer)
		return
	}
}
