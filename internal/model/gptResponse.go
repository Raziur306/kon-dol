package model

type GPTResponse struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	District string `json:"district"`
	Party    struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"party"`
	Date      string `json:"date"`
	ShortDesc string `json:"short_desc"`
	Status    string `json:"status"`
}
