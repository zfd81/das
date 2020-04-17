package meta

type Catalog struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Order  int    `json:"ord"`
	Parent string `json:"parent"`
	Model
}

func (c *Catalog) TableName() string {
	return "das_catalog"
}
