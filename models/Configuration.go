package models

type Auth struct {
	Secret   string `json:"secret"`
	Email    string `json: "email"`
	Password string `json:"password"`
}

type Configuration struct {
	Addr     string `json: "addr"`
	MongoURL string `json: "mongo"`
	SaveDir  string `json: "saveDir"`
	BaseURL  string `json: "baseURL"`
	Auth     Auth   `json: "auth"`
}
