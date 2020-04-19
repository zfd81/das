package dao

import "github.com/zfd81/rooster/types/container"

type ProjectDao interface {
	FindByCode(name string) (*ProjectInfo, error)
	Save(entity *ProjectInfo) error
	FindAllByUser(userId string, codeOrName string) ([]container.Map, error)
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

func (u *ProjectDaoImpl) FindAllByUser(userId string, codeOrName string) ([]container.Map, error) {
	mp := map[string]interface{}{
		"userId":     userId,
		"codeOrName": "%" + codeOrName + "%",
	}
	sql := `
		SELECT * FROM
			(
				SELECT
					p.code,
					p.name,
					u.name creator,
					p.created_time,
					'owner' identity
				FROM
					project p,
					sys_user u
				WHERE
					p.creator = u.id
				AND p.status = '1'
				AND p.creator = :userId
				AND (p.code like :codeOrName OR p.name like :codeOrName)
				UNION
				SELECT
					p.code,
					p.name,
					u.name creator,
					p.created_time,
					'user' identity
				FROM
					project p,
					sys_user u,
					rel_user_project up
				WHERE
					p.creator = u.id
				AND p.code = up.project_code
				AND p.status = '1'
				AND up.user_id = :userId
				AND (up.project_code like :codeOrName OR p.name like :codeOrName)
			) t ORDER BY t.created_time DESC
		`
	r, err := db.Query(sql, mp)
	if err != nil {
		return make([]container.Map, 0, 10), err
	}
	return r.MapListScan()
}

func NewProjectDao() *ProjectDaoImpl {
	return &ProjectDaoImpl{}
}
