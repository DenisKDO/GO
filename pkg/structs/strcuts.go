package structs

// Player struct
type Player struct {
	UniformNumber string `json:"uniformnumber"`
	FirstName     string `json:"firstname"`
	SecondName    string `json:"secondname"`
	Position      string `json:"position"`
	Country       string `json:"country"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
}

// Volleyball team struct
type Response struct {
	Team []Player `json:"team"`
}
