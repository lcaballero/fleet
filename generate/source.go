package generate
import (
	"io/ioutil"
	"html/template"
)

type Source struct {
	Config      *Config
	Source      string
	Dest        string
	Name        string
	Data        interface{}
	rawTemplate string
}

func (src *Source) RawTemplate() (string, error) {
	if src.rawTemplate != "" {
		return src.rawTemplate, nil
	}
	s, err := ioutil.ReadFile(src.Source)
	if err != nil {
		return "", err
	}
	src.rawTemplate = string(s)
	return src.rawTemplate, nil
}

func (src *Source) ToTemplate() (*template.Template, error) {
	tpl, err := src.RawTemplate()
	if err != nil {
		return nil, err
	}

	tt := template.New(src.Name)
	t, err := tt.Parse(tpl)
	if err != nil {
		return nil, err
	}
	return t, nil
}