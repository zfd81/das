package meta

type Service struct {
	Code    string  `json:"code"`
	Name    string  `json:"name"`
	Catalog Catalog `json:"catalog"`
	Method  Method  `json:"method"`
	Params  []Param `json:"params"`
	Model
}

func (s Service) TableName() string {
	return "das_service"
}
