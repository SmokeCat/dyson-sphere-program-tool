package gamedata


type Blueprint struct {
	Time    float32   `json:"time"`
	Product map[string]int  `json:"product`
	Cost    map[string]int  `json:"cost"`
}