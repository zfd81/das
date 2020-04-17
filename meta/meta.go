package meta

import "time"

type Method int

const (
	METHOD_GET Method = iota
	METHOD_POST
	METHOD_PUT
	METHOD_PATCH
	METHOD_DELETE
)

type Model struct {
	Creator      string    `json:"creator"`
	CreatedTime  time.Time `json:"created_time"`
	Modifier     string    `json:"modifier"`
	ModifiedTime time.Time `json:"modified_time"`
}
