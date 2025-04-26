package entity

type Repo struct {
	Name string `json:"name"`
}

type Activity struct {
	Type      string `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
}
