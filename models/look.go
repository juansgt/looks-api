package models

type Look struct {
	id    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func NewLook(id string, name string, brand string) Look {
	return Look{
		id:    id,
		Name:  name,
		Brand: brand,
	}
}

func (look *Look) Id() string {
	return look.id
}
