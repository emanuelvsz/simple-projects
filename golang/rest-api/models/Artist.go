package models

type Artist struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Genre []string `json:"genres"`
	Age   int      `json:"age"`
	Songs []string `json:"songs"`
}
