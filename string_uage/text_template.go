package stringuage

import (
	"html/template"
	"os"
)

func Test_text_template() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	//t1 := template.Must(t1.Parse("Value is {{.}}\n"))  // we can use the template.Must function to panic in case Parse returns an error

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, "5")
	t1.Execute(os.Stdout, []string{"Go", "Rust"})

	Create := func(name string, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	} // a convinent tool to creat a template

	t2 := Create("t2", "Name: {{.Name}}\n") // get value from struct or map
	t2.Execute(os.Stdout, struct {
		Name string
	}{"JanessaTech"}) // define anoymous struct
	t2.Execute(os.Stdout, map[string]string{"Name": "this is value from map"})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

}
