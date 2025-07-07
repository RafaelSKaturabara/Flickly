package entities

type Action struct {
	Prioridade     int
	Condition      *Condition
	Type           string
	FieldToSet     string
	CookieName     string
	CookieValue    string
	RemoveCookie   bool
	Result         string
}
