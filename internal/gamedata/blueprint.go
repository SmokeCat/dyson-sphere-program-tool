package gamedata


type Product struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}
type Cost struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}
type Blueprint struct {
	Id      int       `json:"id"`
	Time    float32   `json:"time"`
	Product []Product `json:"product"`
	Cost    []Cost    `json:"cost"`
}