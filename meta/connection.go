package meta

type Connection struct {
	DriverName   string `json:"driver"`
	Address      string `json:"address"`
	Port         int    `json:"port"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	DatabaseName string `json:"db"`
	Project      *Project
}
