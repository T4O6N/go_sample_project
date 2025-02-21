package entities

type Subject struct {
	ID      int    `json:"id"`
	Name_en string `json:"name_en"`
	Name_la string `json:"name_la"`
	Status  bool   `json:"status"`
}
