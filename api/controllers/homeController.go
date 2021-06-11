package controllers

import (
	"net/http"

	"levelup/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API kuliner")

}