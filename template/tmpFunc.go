package template

import (
	"bytes"
	"strings"
	"text/template"
)

/*
格式 id int64,fal_id int64

stu
id int64
fal_id int64
name string


*/

func GetTemplate(tmpl string) (*template.Template, error) {
	t := template.New("etool")
	t.Funcs(template.FuncMap{
		"FieldName": FieldName,
		"Jsondata":  Jsondata,
	})

	return t.Parse(tmpl)
}

//转编译后的字符串
func ExecTemplate(tpl *template.Template, data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := tpl.Execute(buf, data)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

//field name
func FieldName(s string) string {
	s = replaceInvalidChars(strings.TrimSpace(s))

	strlist := strings.Split(s, "_")
	out := ""
	for _, v := range strlist {
		out += UpFirstChar(v)
	}
	for k, list := range uPerStr {
		for _, v := range list {
			out = strings.ReplaceAll(out, v, k)
		}

	}

	return out
}

//fieldName to json name
func Jsondata(s string) string {
	s = FieldName(s)
	strlist := strings.Split(s, "_")
	out := ""
	for _, v := range strlist {
		if isWordInUPword(s) {
			out += UpFirstChar(v)
		} else {
			out += LowFirstChar(v)
		}

	}
	return out
}

var uPerStr = map[string][]string{
	"ID": []string{"Id", "ID"},
}
