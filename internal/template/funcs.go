package template

import (
	"encoding/json"
	"os"
	"strings"
	"swan/internal/memkv"
	"time"
)

//newFuncMap 自定一个模板渲染函数
func newFuncMap() map[string]interface{} {
	fnMap := memkv.GetKv().FuncMap()

	fnMap["join"] = strings.Join
	fnMap["trim"] = strings.Trim
	fnMap["json"] = jsonUnmarshal
	fnMap["title"] = strings.Title
	fnMap["index"] = strings.Index
	fnMap["count"] = strings.Count
	fnMap["split"] = strings.Split
	fnMap["repeat"] = strings.Repeat
	fnMap["splitn"] = strings.SplitN
	fnMap["getenv"] = os.Getenv
	fnMap["replace"] = strings.Replace
	fnMap["tolower"] = strings.ToLower
	fnMap["toupper"] = strings.ToUpper
	fnMap["contains"] = strings.Contains
	fnMap["trimleft"] = strings.TrimLeft
	fnMap["strftime"] = strftime
	fnMap["hasprefix"] = strings.HasPrefix
	fnMap["hassuffix"] = strings.HasSuffix
	fnMap["trimright"] = strings.TrimRight
	fnMap["trimspace"] = strings.TrimSpace
	fnMap["equalfold"] = strings.EqualFold
	fnMap["trimprefix"] = strings.TrimPrefix
	fnMap["trimsuffix"] = strings.TrimSuffix
	fnMap["replaceall"] = strings.ReplaceAll

	return fnMap
}

func jsonUnmarshal(s string) (v interface{}, err error) {
	err = json.Unmarshal([]byte(s), &v)
	return
}

func strftime(value, layout, fmt string) (string, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return "", err
	}
	return t.Format(fmt), nil
}
