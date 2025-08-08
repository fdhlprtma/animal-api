package model

type Classification struct {
	Kingdom string `json:"kingdom"`
	Phylum  string `json:"phylum"`
	Class   string `json:"class"`
	Order   string `json:"order"`
	Family  string `json:"family"`
}

type Animal struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	ImageURL        string         `json:"image_url"`
	Classification  Classification `json:"classification"`
	Characteristics string         `json:"characteristics"`
	Examples        []string       `json:"examples"`
	Habitat         string         `json:"habitat"`
	EcologicalRole  string         `json:"ecological_role"`
}
