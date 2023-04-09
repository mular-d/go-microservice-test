package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPaylod mailMessage

	err := app.readJSON(w, r, &requestPaylod)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPaylod.From,
		To:      requestPaylod.To,
		Subject: requestPaylod.Subject,
		Data:    requestPaylod.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPaylod.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
