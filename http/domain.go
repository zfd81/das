package http

type TreeNode struct {
	Id       string     `json:"id"`
	Label    string     `json:"label"`
	Creator  string     `json:"creator"`
	Type     string     `json:"type"`
	Children []TreeNode `json:"children"`
}
