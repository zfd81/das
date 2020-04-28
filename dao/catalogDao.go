package dao

import "github.com/zfd81/rooster/types/container"

type CatalogDao interface {
	Save(entity *CatalogInfo) error
	DeleteByCode(code string) error
	Modify(entity *CatalogInfo) error
	ModifyStatus(status string, code string) error
	FindByCode(code string) (container.Map, error)
	FindAllByParent(parentCode string) ([]container.Map, error)
}

type CatalogDaoImpl struct {
}

func (c *CatalogDaoImpl) Save(entity *CatalogInfo) (err error) {
	_, err = db.Save(entity)
	return
}

func (c *CatalogDaoImpl) DeleteByCode(code string) (err error) {
	sql := "delete from das_catalog where catalog_code=:val"
	_, err = db.Exec(sql, code)
	return
}

func (c *CatalogDaoImpl) Modify(entity *CatalogInfo) (err error) {
	sql := `
		UPDATE das_catalog SET 
  			catalog_name =: catalog_name,
		 	modifier =: modifier,
			modified_time =: modified_time
		WHERE
			catalog_code =: catalog_code
	`
	_, err = db.Exec(sql, entity)
	return
}

func (c *CatalogDaoImpl) ModifyStatus(status string, code string) (err error) {
	catalog := &CatalogInfo{Code: code, Status: status}
	sql := `
		UPDATE das_catalog SET 
  			status =: status,
		 	modifier =: modifier,
			modified_time =: modified_time
		WHERE
			catalog_code =: catalog_code
	`
	_, err = db.Exec(sql, catalog)
	return
}

func (c *CatalogDaoImpl) FindByCode(code string) (container.Map, error) {
	sql := "select catalog_code,catalog_name,ord from das_catalog where catalog_code=:val"
	return db.QueryMap(sql, code)
}

func (c *CatalogDaoImpl) FindAllByParent(parentCode string) ([]container.Map, error) {
	sql := "select catalog_code,catalog_name,ord,parent_code,creator from das_catalog where parent_code=:val ORDER BY catalog_name"
	r, err := db.Query(sql, parentCode)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func NewCatalogDao() *CatalogDaoImpl {
	return &CatalogDaoImpl{}
}
