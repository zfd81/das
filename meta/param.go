package meta

import "github.com/zfd81/das/types"

type Param struct {
	Code         string      `json:"code"`
	Name         string      `json:"name"`
	Type         types.Type  `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
}
