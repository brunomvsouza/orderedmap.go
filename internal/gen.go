package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	naturalType := flag.String("naturalType", "", "the natural type of the ordered map keys")
	noTest := flag.Bool("noTest", false, "do not generate tests")
	flag.Parse()

	if *naturalType == "" {
		panic("naturalType can't be empty")
	}

	data := struct {
		NaturalType         string
		TypeName            string
		TypeNameCapitalized string
		Generic             bool
	}{}

	data.NaturalType = *naturalType
	data.TypeName = func() string {
		if *naturalType == "interface{}" {
			return "GenericMap"
		}

		typeName := *naturalType
		if strings.HasPrefix(typeName, "*") {
			typeName = strings.Replace(fmt.Sprintf("%sPtr", typeName), "*", "", 1)
		}
		if strings.HasPrefix(typeName, "[") {
			tmp := strings.Replace(typeName, "[", "", 1)
			tmp = strings.Replace(tmp, "]byte", "", 1)
			typeName = fmt.Sprintf("byte%s", tmp)
		}
		typeName = fmt.Sprintf("%sMap", typeName)
		return strings.Title(typeName)
	}()
	data.Generic = data.TypeName == "GenericMap"

	f, err := os.Create(fmt.Sprintf("zgenerated_%s.go", strings.ToLower(data.TypeName)))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tpl := template.Must(template.ParseFiles("internal/gen.tpl"))
	err = tpl.Execute(f, data)
	if err != nil {
		panic(err)
	}

	if *noTest {
		return
	}

	ft, err := os.Create(fmt.Sprintf("zgenerated_%s_test.go", strings.ToLower(data.TypeName)))
	if err != nil {
		panic(err)
	}
	defer ft.Close()

	tplt := template.Must(template.ParseFiles("internal/gen_test.tpl"))
	err = tplt.Execute(ft, data)
	if err != nil {
		panic(err)
	}
}
