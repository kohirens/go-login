package login

type SubProfile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (sp *SubProfile) String() string {
	jsonString := `"id": "` + sp.Id + `",`
	jsonString += `"name": "` + sp.Name + `"`
	return "{" + jsonString + "}"
}
