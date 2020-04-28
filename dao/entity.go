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
	return "das_sys_user"
}

type ProjectInfo struct {
	Code        string `rsql:"name:code"`
	Name        string `rsql:"name:name"`
	Description string `rsql:"name:description"`
	Status      string `rsql:"name:status"`
	Model
}

func (p ProjectInfo) TableName() string {
	return "das_project"
}

type CatalogInfo struct {
	Code    string `rsql:"name:catalog_code"`
	Name    string `rsql:"name:catalog_name"`
	Order   int    `rsql:"name:ord"`
	Parent  string `rsql:"name:parent_code"`
	Project string `rsql:"name:project_code"`
	Status  string `rsql:"name:status"`
	Model
}

func (c CatalogInfo) TableName() string {
	return "das_catalog"
}

type ConnectionInfo struct {
	ID           string `rsql:"name:conn_id"`
	Name         string `rsql:"name:conn_name"`
	Driver       string `rsql:"name:driver"`
	Address      string `rsql:"name:address"`
	Port         string `rsql:"name:port"`
	UserName     string `rsql:"name:user_name"`
	Password     string `rsql:"name:password"`
	DatabaseName string `rsql:"name:db"`
	Project      string `rsql:"name:project_code"`
	Status       string `rsql:"name:status"`
	Model
}

func (c ConnectionInfo) TableName() string {
	return "das_connection"
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
