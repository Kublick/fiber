package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

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
	ID              string `json:"id"`
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

func getExpedientById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range expedients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}

var jwtKey = []byte("Mysuperpassword")
var api_key = "1234"

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return jwtKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func GetJwt(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Logging")
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

var allowedUsers = map[string]string{
	"tester":  "Testing01!",
	"tester2": "Testing01!",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Entered")
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := allowedUsers[creds.Username]
	if (!ok) || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expiration := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if (err) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiration,
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
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
		ID:              "1",
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
		ID:              "2",
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
		ID:              "6",
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

	r.HandleFunc("/", GetJwt).Methods("GET")
	r.HandleFunc("/api", Home).Methods("POST")
	r.HandleFunc("/tiers", getTiers).Methods("GET")
	r.HandleFunc("/workspaces", getWorkspaces).Methods("GET")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/expedients", getExpedients).Methods("GET")
	r.HandleFunc("/expedients/{id}", getExpedientById).Methods("GET")

	log.Fatal(http.ListenAndServe(getPort(), r))
}
