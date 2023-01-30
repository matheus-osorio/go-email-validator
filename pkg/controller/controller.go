package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheus-osorio/go-email-validator/pkg/model/database"
	"github.com/matheus-osorio/go-email-validator/pkg/model/sender"
	"github.com/matheus-osorio/go-email-validator/pkg/model/verifier"
	"github.com/matheus-osorio/go-email-validator/pkg/utils"
)

func CheckDomainExistence(writer http.ResponseWriter, request *http.Request) {
	domain := request.Header.Get("domain")
	response := verifier.VerifyEmail(domain)

	parametizer := utils.CreateParametizer()
	parametizer.Headers.SetHeader("Content-Type", "application/json")
	parametizer.SetStatus(http.StatusOK)

	data, _ := json.Marshal(response)
	parametizer.SetBody(data)
	parametizer.Respond(writer)

}

func CheckOwnershipEmail(writer http.ResponseWriter, request *http.Request) {
	var email database.Email
	utils.ParseBody(request, &email)
	email.Create()
	sender.SendVerification(email.Email)
	parametizer := utils.CreateParametizer()
	parametizer.Headers.SetHeader("Content-Type", "application/json")
	parametizer.SetStatus(http.StatusOK)

	body := struct {
		Message string
		Uuid    string
	}{
		Message: "An email has been sent to the account " + email.Email,
		Uuid:    email.UUID,
	}
	data, _ := json.Marshal(body)
	parametizer.SetBody(data)
	parametizer.Respond(writer)
}

func ReceiveOwnershipEmail(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	uuid := vars["uuid"]
	email := database.Email{
		UUID: uuid,
	}
	email.ReadSingle()
	if email.ID != 0 {
		email.Validated = true
		email.Update()
	}

	parametizer := utils.CreateParametizer()
	parametizer.Headers.SetHeader("Content-Type", "application/json")
	parametizer.SetStatus(http.StatusOK)

	body := struct {
		Message string
		Uuid    string
	}{
		Message: "Email  " + email.Email + " has been validated successfully",
		Uuid:    email.UUID,
	}
	data, _ := json.Marshal(body)
	parametizer.SetBody(data)
	parametizer.Respond(writer)
}
