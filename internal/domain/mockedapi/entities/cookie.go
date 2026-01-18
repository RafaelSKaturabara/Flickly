package entities

type Cookie struct {
	Name     string
	Value    string
	Domain   string
	Path     string
	Expires  int64
	Secure   bool
	HttpOnly bool
}
