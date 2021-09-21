package model

// Template is used to pass the data to create a template
type Template struct {
	Name        string
	Filename    string
	Path        string
	WithPackage bool
	Layer       Layer
}

func (t *Template) SetPathPrefix(data Layer) {
	t.Path = data.GetPath(t.Path, t.Filename, t.WithPackage)
}

func (t *Template) SetLayerData(data Layer) {
	t.Layer = data
}

// Templates slice of Template
type Templates []Template

// SetPath is used when the path is the same for every template
func (t Templates) SetPath(path string) {
	for _, v := range t {
		v.Path = path
	}
}
