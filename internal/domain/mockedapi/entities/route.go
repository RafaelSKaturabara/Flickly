package entities


type Route struct {
	Paths   string
	Methods HttpMethod
	Mocks   []*Mock
}

func NewRoute(paths string, methods HttpMethod, mocks []*Mock) *Route {
	return &Route{
		Paths:   paths,
		Methods: methods,
		Mocks:   mocks,
	}
}