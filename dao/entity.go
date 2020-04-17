package dao

import "time"

type Model struct {
	Creator      string    `rsql:"name:creator"`
	CreatedTime  time.Time `rsql:"name:created_time"`
	Modifier     string    `rsql:"name:modifier"`
	ModifiedTime time.Time `rsql:"name:modified_time"`
}

type UserInfo struct {
	ID          string `rsql:"name:id"`
	Name        string `rsql:"name:name"`
	Password    string `rsql:"name:password"`
	FullName    string `rsql:"name:full_name"`
	PhoneNumber string `rsql:"name:phone_number"`
	Email       string `rsql:"name:email"`
	Status      string `rsql:"name:status"`
	Model
}

func (u UserInfo) TableName() string {
	return "sys_user"
}

type ProjectInfo struct {
	Code        string `rsql:"name:code"`
	Name        string `rsql:"name:name"`
	Description string `rsql:"name:description"`
	Status      string `rsql:"name:status"`
	Model
}

func (p ProjectInfo) TableName() string {
	return "project"
}

type Catalog struct {
	Code   string `rsql:"name:code"`
	Name   string `rsql:"name:name"`
	order  int    `rsql:"name:ord"`
	Parent string `rsql:"name:parent_id"`
	Model
}

func (c Catalog) TableName() string {
	return "das_catalog"
}

type Service struct {
	Code    string `rsql:"name:code"`
	Name    string `rsql:"name:name"`
	Method  int    `rsql:"name:method"`
	Catalog string `rsql:"name:catalog_id"`
	Model
}

func (s Service) TableName() string {
	return "das_service"
}
