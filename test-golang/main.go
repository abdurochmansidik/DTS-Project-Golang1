package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pandawa struct {
	Nama         string `json:"nama"`
	Julukan      string `json:"nama_julukan"`
	Titisan      string `json:"titisan"`
	DiangkatRaja bool   `json:"raja"`
}

func (p *Pandawa) diAngkatRaja() {
	p.DiangkatRaja = true
}

func main() {
	pandawas := []Pandawa{
		{
			Nama:    "Yudistira",
			Julukan: "Dhramasuta",
			Titisan: "Dewa Yama",
		},
		{
			Nama:    "Bima",
			Julukan: "Bayusutha",
			Titisan: "Dewa Bayu",
		},
		{
			Nama:    "Arjuna",
			Julukan: "Partha",
			Titisan: "Dewa Indra",
		},
		{
			Nama:    "Nakula",
			Julukan: "pengasuh kuda",
			Titisan: "Dewa Aswin",
		},
		{
			Nama:    "Sadewa",
			Julukan: "Brihaspati",
			Titisan: "Dewa Aswin",
		},
	}

	pandawas[0].diAngkatRaja()

	jsonBytes, err := json.Marshal(pandawas)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonBytes))
}
