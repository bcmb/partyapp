package docs

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"io/ioutil"
	"net/http"
)

// @Summary      Get all upcoming parties
// @Description  List parties
// @Tags         party
// @Accept       json
// @Produce      json
// @Security BasicAuth
// @Success      200  {object}  party.Party
// @Router       /party [get]
func getAllParties() {}

// @Summary      Get one party
// @Description  Get a single party by id
// @Tags         party
// @Accept       json
// @Produce      json
// @Security BasicAuth
// @Param        id   path      int  true  "party Id"
// @Success      200  {object}  party.Party
// @Router       /party/{id} [get]
func getOneParty() {}

// @Summary      Create a new party
// @Description  Creates new party
// @Tags         party
// @Accept       json
// @Produce      json
// @Security BasicAuth
// @Param request body party.Party true "Party Data"
// @Success      200  {object}  party.Party
// @Router       /party [post]
func createParty() {}

// @Summary      Update existing party
// @Description  updates existing party
// @Tags         party
// @Accept       json
// @Produce      json
// @Security BasicAuth
// @Param        id   path      int  true  "party Id"
// @Param request body party.Party true "Party Data"
// @Success      200  {object}  party.Party
// @Router       /party/{id} [patch]
func updateParty() {}

// @Summary      Delete existing party
// @Description  deletes party
// @Tags         party
// @Accept       json
// @Produce      json
// @Security BasicAuth
// @Param        id   path      int  true  "party Id"
// @Success      200  {object}  party.Party
// @Router       /party/{id} [delete]
func deleteParty() {}

func SetupRoutes() {
	http.Handle(fmt.Sprintf("/swagger/"), httpSwagger.Handler())
	docs, _ := ioutil.ReadFile("docs/swagger.json")
	http.HandleFunc(fmt.Sprintf("/swagger/doc.json"), func(w http.ResponseWriter, r *http.Request) { w.Write(docs) })
}
