package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPaylod JSONPayload
	_ = app.readJSON(w, r, &requestPaylod)

	event := data.LogEntry{
		Name: requestPaylod.Name,
		Data: requestPaylod.Data,
	}

	err := app.Modles.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
