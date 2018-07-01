package config

import (
	"encoding/json"

	"os"
)

type confguration struct {
	Blist []string `json:"backlist"`
}

func LoadConfig() (s []string) {

	file, _ := os.Open("./config.json")
	defer file.Close()
	docoder := json.NewDecoder(file)
	stu := confguration{}
	docoder.Decode(&stu)
	return stu.Blist
}
