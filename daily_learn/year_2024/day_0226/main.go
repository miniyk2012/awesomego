package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type BBQBuiltin struct {
	A string
}

func ResolveTemplate(source string, paramMap map[string]any, funcMap template.FuncMap) (string, error) {
	t := template.New("bbqExpression")
	if funcMap != nil {
		t = t.Funcs(funcMap)
	}
	parsed, err := t.Parse(source)
	if err != nil {
		return source, err
	}
	buffer := bytes.NewBuffer(nil)
	err = parsed.Execute(buffer, paramMap)
	if err != nil {
		return source, err
	}
	return buffer.String(), nil
}

const source = `{{bbq.A}}`

var bbqBuiltin = BBQBuiltin{
	A: "hello",
}

type User struct {
	ID    int
	Email string
}
type ViewData struct {
	User User
}

func example() {
	bbqFuncMap := map[string]any{
		"bbq": func() BBQBuiltin { return bbqBuiltin },
	}
	v, err := ResolveTemplate(source, nil, bbqFuncMap)
	fmt.Println(err)
	fmt.Printf("%v", v)
}

func another() {
	var err error
	testTemplate, err := template.New("hello.gohtml").Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			return false
		},
	}).ParseFiles("daily_learn/year_2024/day_0226/hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		ID:    1,
		Email: "jon@calhoun.io",
	}
	//vd := ViewData{user}
	//testTemplate.Execute(os.Stdout, vd)
	err = template.Must(testTemplate.Clone()).Funcs(template.FuncMap{
		"hasPermission": func(feature string) bool {
			if user.ID == 1 && feature == "feature-a" {
				return true
			}
			return false
		},
	}).Execute(os.Stdout, nil)
	fmt.Println(err)
}

func main() {
	//example()
	another()
}
