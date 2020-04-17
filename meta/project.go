package meta

type Project struct {
	Code     string       `json:"code"`
	Name     string       `json:"name"`
	Conns    []Connection `json:"conns"`
	Catalogs []Catalog    `json:"catalogs"`
}
