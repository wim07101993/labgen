package svg

type Id string

func (id Id) Href() string {
	return "#" + string(id)
}
