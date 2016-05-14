package backup

type Element struct {
	Name     string
	Path     string
	Excludes []string
}

func (u Element) IsValid() bool {
	return u.Name != ""
}
