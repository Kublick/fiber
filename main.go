package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Tier struct {
	ID                        int64 `json:"id"`
	UsersAmount               int64 `json:"users_amount"`
	CompanionApp              bool  `json:"companion_app"`
	UpdateNotifications       bool  `json:"update_notifications"`
	LogDuration               int64 `json:"log_duration"`
	Storage                   int64 `json:"storage"`
	ClientsEmailNotifications bool  `json:"clients_email_notifications"`
	CostControl               bool  `json:"cost_control"`
	Price                     int64 `json:"price"`
}

var tiers []Tier

type Workspace struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Alias               string `json:"alias"`
	LegalRepresentative string `json:"legalRepresentative"`
	Picture             string `json:"picture"`
	Phone               string `json:"phone"`
	TierID              int64  `json:"tierId"`
}

var workspaces []Workspace

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Carrier  string `json:"carrier"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
	ID       int64  `json:"id"`
}

var users []User

type Expedient struct {
	ID              int64  `json:"id"`
	WorkspaceID     int64  `json:"workspaceId"`
	InternalID      int64  `json:"internalId"`
	ClientName      string `json:"clientName"`
	ClientEmail     string `json:"clientEmail"`
	ClientPhone     string `json:"clientPhone"`
	ExpedientNumber string `json:"expedientNumber"`
	GroupNumber     string `json:"groupNumber"`
	UserID          int64  `json:"userId"`
	Category        string `json:"category"`
	Counterpart     string `json:"counterpart"`
	Status          string `json:"status"`
	AuthorityID     int64  `json:"authorityId"`
	Book            int64  `json:"book"`
	Amparo          int64  `json:"amparo"`
	Quantity        int64  `json:"quantity"`
	Currency        string `json:"currency"`
	InitialDate     string `json:"initialDate"`
}

var expedients []Expedient

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func getTiers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tiers)

}

func getWorkspaces(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workspaces)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getExpedients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expedients)
}

func main() {

	tiers = append(tiers, Tier{
		ID:                        1,
		UsersAmount:               3,
		CompanionApp:              true,
		UpdateNotifications:       true,
		LogDuration:               15,
		Storage:                   5,
		ClientsEmailNotifications: false,
		CostControl:               false,
		Price:                     10000,
	})
	tiers = append(tiers, Tier{
		ID:                        2,
		UsersAmount:               10,
		CompanionApp:              true,
		UpdateNotifications:       true,
		LogDuration:               60,
		Storage:                   10,
		ClientsEmailNotifications: false,
		CostControl:               false,
		Price:                     20000,
	})

	workspaces = append(workspaces, Workspace{
		ID:                  1,
		Name:                "Franco y Asociados",
		Alias:               "francoya",
		LegalRepresentative: "Juan Leal Franco",
		Picture:             "https://source.unsplash.com/user/c_v_r/100x100",
		Phone:               "6862424242",
		TierID:              1,
	})
	workspaces = append(workspaces, Workspace{
		ID:                  2,
		Name:                "Lopez Abogados",
		Alias:               "LopezAbogados",
		LegalRepresentative: "Juan Leal Franco",
		Picture:             "https://source.unsplash.com/user/c_v_r/100x100",
		Phone:               "6862424242",
		TierID:              1,
	})
	workspaces = append(workspaces, Workspace{
		ID:                  3,
		Name:                "Abshire-Satterfield",
		Alias:               "gdocharty0",
		LegalRepresentative: "Gavra Docharty",
		Picture:             "https://source.unsplash.com/user/c_v_r/100x100",
		Phone:               "180-682-1953",
		TierID:              2,
	})
	workspaces = append(workspaces, Workspace{
		ID:                  4,
		Name:                "Fisher, Schmidt and Roob",
		Alias:               "rreynault2",
		LegalRepresentative: "Romeo Reynault",
		Picture:             "http://localhost:3000/admin/register/workspace/Builder.png",
		Phone:               "588-362-6556",
		TierID:              3,
	})

	users = append(users, User{
		Name:     "Dahlia Gillespie",
		Email:    "test@test.com",
		Phone:    "6862443020",
		Carrier:  "telcel",
		Password: "Testing01!",
		Picture:  "https://i.pravatar.cc/100",
		ID:       1,
	})
	users = append(users, User{
		Name:     "Harry Mason",
		Email:    "test2@test.com",
		Phone:    "6862443020",
		Carrier:  "telcel",
		Password: "Testing01!",
		Picture:  "https://i.pravatar.cc/100",
		ID:       2,
	})

	expedients = append(expedients, Expedient{
		ID:              1,
		WorkspaceID:     1,
		InternalID:      1,
		ClientName:      "Pedro Vargas",
		ClientEmail:     "pvargas@cielo.com",
		ClientPhone:     "5553241569",
		ExpedientNumber: "30102021AB",
		GroupNumber:     "215",
		UserID:          1,
		Category:        "familiar",
		Counterpart:     "Jorge Aguilar",
		Status:          "En espera de acuerdo",
		AuthorityID:     6,
		Book:            4,
		Amparo:          21321313,
		Quantity:        5,
		Currency:        "mxn",
		InitialDate:     "2022-09-09",
	})

	expedients = append(expedients, Expedient{
		ID:              1,
		WorkspaceID:     2,
		InternalID:      2,
		ClientName:      "Jose Alfredo Jimenez",
		ClientEmail:     "JJimenez@cielo.com",
		ClientPhone:     "5636581110",
		ExpedientNumber: "33112441CD",
		GroupNumber:     "29",
		UserID:          1,
		Category:        "familiar",
		Counterpart:     "Laureano Brizuela",
		Status:          "En espera de acuerdo",
		AuthorityID:     7,
		Book:            6,
		Amparo:          11111111,
		Quantity:        100,
		Currency:        "mxsn",
		InitialDate:     "2022-09-09",
	})
	expedients = append(expedients, Expedient{
		ID:              6,
		WorkspaceID:     7,
		InternalID:      2,
		ClientName:      "Armando Cruz",
		ClientEmail:     "cruzarmando@cielo.com",
		ClientPhone:     "1234567891",
		ExpedientNumber: "32323213",
		GroupNumber:     "29",
		UserID:          1,
		Category:        "civil",
		Counterpart:     "Jacki Allcroft",
		Status:          "En espera de acuerdo",
		AuthorityID:     7,
		Book:            6,
		Amparo:          11111111,
		Quantity:        100,
		Currency:        "mxsn",
		InitialDate:     "2022-09-09",
	})

	r := mux.NewRouter()

	r.HandleFunc("/tiers", getTiers).Methods("GET")
	r.HandleFunc("/workspaces", getWorkspaces).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/expedients", getExpedients).Methods("GET")

	log.Fatal(http.ListenAndServe(getPort(), r))
}
