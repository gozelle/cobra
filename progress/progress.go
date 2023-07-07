package progress

import "fmt"

func Value(val string) *Field {
	return &Field{
		Value: val,
	}
}

type Field struct {
	Width uint
	Value string
}

func (p *Field) WithWidth(l uint) *Field {
	p.Width = l
	return p
}

func (p *Field) Format() string {
	if p.Width == 0 {
		return "%s"
	}
	return fmt.Sprintf("%%-%ds", p.Width)
}
