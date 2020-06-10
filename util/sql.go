package util

import (
	"strings"

	"github.com/zfd81/rooster/util"
)

type KV struct {
	Key string
	Val interface{}
}

func validCharacter(char byte) bool {
	if (char >= 48 && char <= 57) || (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || char == 46 || char == 95 || char == 9 || char == 10 {
		return true
	}
	return false
}

func ParsingSql(sql string) ([]KV, error) {
	params := make([]KV, 0, 20)
	start := 0 //切片名称开始位置
	end := 0   //切片名称结束位置
	open := 0  //{大括号开始位置
	close := 0 //}大括号结束位置
	for i, char := range sql {
		if char == '{' && open == 0 {
			open = i + 1
			end = i
		} else if char == '}' && open > 0 {
			close = i
			_, err := util.ReplaceByKeyword(sql[start:end], ':', func(index int, start int, end int, content string) (string, error) {
				params = append(params, KV{content, content})
				return "?", nil
			})
			if err != nil {
				return nil, err
			}

			s := 0
			e := 0
			flag := false
			content := sql[open:close]
			for i, char := range content {
				if flag {
					if char == 32 && s == 0 {
						continue
					} else if s == 0 && validCharacter(byte(char)) {
						s = i
						continue
					} else if !validCharacter(byte(char)) && s > 0 {
						e = i
						break
					}
				} else {
					if char == '@' && s == 0 && e == 0 {
						flag = true
					}
				}
			}
			if flag {
				name := strings.TrimSpace(content[s:e])
				ps := make([]KV, 0, 20)
				_, err := util.ReplaceByKeyword(content, ':', func(index int, start int, end int, content string) (string, error) {
					if strings.HasPrefix(content, "this.") {
						key := content[5:]
						ps = append(ps, KV{key, key})
					} else {
						params = append(params, KV{content, content})
					}
					return "?", nil
				})
				if err != nil {
					return nil, err
				}
				if len(ps) > 0 {
					params = append(params, KV{name, ps})
				}
			} else {
				_, err := util.ReplaceByKeyword(sql[start:], ':', func(index int, start int, end int, content string) (string, error) {
					params = append(params, KV{content, content})
					return "?", nil
				})
				if err != nil {
					return nil, err
				}
			}
			start = close + 1
			end = start
			open = 0
			close = 0
		}
	}
	_, err := util.ReplaceByKeyword(sql[start:], ':', func(index int, start int, end int, content string) (string, error) {
		params = append(params, KV{content, content})
		return "?", nil
	})
	if err != nil {
		return nil, err
	}
	return params, nil
}

//func ParsingSql(sql string) ([]KV, error) {
//	params := make([]KV, 0, 20)
//	start := 0 //切片名称开始位置
//	end := 0   //切片名称结束位置
//	open := 0  //{大括号开始位置
//	close := 0 //}大括号结束位置
//	for i, char := range sql {
//		if char == '{' && open == 0 {
//			open = i + 1
//			end = i
//		} else if char == '}' && open > 0 {
//			close = i
//			_, err := util.ReplaceByKeyword(sql[start:end], ':', func(index int, start int, end int, content string) (string, error) {
//				params = append(params, KV{content, content})
//				return "?", nil
//			})
//			if err != nil {
//				return nil, err
//			}
//			if sql[open] == '@' {
//				start := 0
//				end := 0
//				content := sql[open:close]
//				for i, char := range content {
//					if char == 32 && start == 0 {
//						continue
//					} else if start == 0 && validCharacter(byte(char)) {
//						start = i
//						continue
//					} else if !validCharacter(byte(char)) && start > 0 {
//						end = i
//						break
//					}
//				}
//				name := strings.TrimSpace(content[start:end])
//				ps := make([]KV, 0, 20)
//				_, err := util.ReplaceByKeyword(content, ':', func(index int, start int, end int, content string) (string, error) {
//					if strings.HasPrefix(content, "this.") {
//						key := content[5:]
//						ps = append(ps, KV{key, key})
//					} else {
//						params = append(params, KV{content, content})
//					}
//					return "?", nil
//				})
//				if err != nil {
//					return nil, err
//				}
//				if len(ps) > 0 {
//					params = append(params, KV{name, ps})
//				}
//			}
//			start = close + 1
//			end = start
//			open = 0
//			close = 0
//		}
//	}
//	_, err := util.ReplaceByKeyword(sql[start:], ':', func(index int, start int, end int, content string) (string, error) {
//		params = append(params, KV{content, content})
//		return "?", nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return params, nil
//}

//
//func ParsingSql(sql string) (map[string]interface{}, error) {
//	params := map[string]interface{}{}
//	newSql, err := util.ReplaceBetween(sql, "{", "}", func(index int, start int, end int, content string) (string, error) {
//		if content != "" {
//			start := 0 //切片名称开始位置
//			end := 0   //切片名称结束位置
//			if content[0] == '@' {
//				for i, char := range content {
//					if char == 32 && start == 0 {
//						continue
//					} else if start == 0 && validCharacter(byte(char)) {
//						start = i
//						continue
//					} else if !validCharacter(byte(char)) && start > 0 {
//						end = i
//						break
//					}
//				}
//				name := strings.TrimSpace(content[start:end])
//				p := map[string]string{}
//				str, err := util.ReplaceByKeyword(content, ':', func(index int, start int, end int, content string) (string, error) {
//					p[content] = content
//					return "?", nil
//				})
//				if err != nil {
//					return "", nil
//				}
//				params[name] = p
//				return str, nil
//			} else {
//				return content, nil
//			}
//		}
//		return "", nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	newSql, err = util.ReplaceByKeyword(newSql, ':', func(index int, start int, end int, content string) (string, error) {
//		params[content] = content
//		return "?", nil
//	})
//	return params, err
//}
