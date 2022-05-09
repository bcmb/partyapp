package party

import (
	"encoding/json"
	"fmt"
	"gotraining.com/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const servicePath = "party"

func handleParties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch r.Method {
	case http.MethodGet:
		parties, err := GetParties()

		if err != nil {
			log.Print(err)
		}

		j, err := json.Marshal(parties)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)

	case http.MethodPost:
		var party *Party

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		err = json.Unmarshal(body, &party)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = InsertParty(party)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleParty(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", servicePath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	partyId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch r.Method {
	case http.MethodGet:
		party, err := GetParty(partyId)
		if err != nil {
			log.Print(err)
		}

		var j []byte
		if party.Id != 0 {
			j, err = json.Marshal(party)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(j)

	case http.MethodPatch:
		var party *Party

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		err = json.Unmarshal(body, &party)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = UpdateParty(party, partyId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		err = DeleteParty(partyId)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes(baseApiPath string) {
	securedPartiesHandler := middleware.BasicAuthMiddleware(handleParties)
	securedPartyHandler := middleware.BasicAuthMiddleware(handleParty)
	http.Handle(fmt.Sprintf("%s/%s", baseApiPath, servicePath), securedPartiesHandler)
	http.Handle(fmt.Sprintf("%s/%s/", baseApiPath, servicePath), securedPartyHandler)
}
