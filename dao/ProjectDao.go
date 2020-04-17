package dao

type ProjectDao interface {
	FindByCode(name string) (*ProjectInfo, error)
	Save(entity *ProjectInfo) error
}

type ProjectDaoImpl struct {
}

func (u *ProjectDaoImpl) FindByCode(code string) (*ProjectInfo, error) {
	project := &ProjectInfo{}
	sql := "select code from project where code=:val"
	err := db.QueryStruct(project, sql, code)
	return project, err
}

func (p *ProjectDaoImpl) Save(entity *ProjectInfo) (err error) {
	_, err = db.Save(entity)
	return
}

func NewProjectDao() *ProjectDaoImpl {
	return &ProjectDaoImpl{}
}
