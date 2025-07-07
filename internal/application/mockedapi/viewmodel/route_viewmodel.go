package viewmodel

type RouteViewModel struct {
	Paths   []string          `json:"paths"`
	Methods []string          `json:"methods"`
	Mocks   []*MockViewModel  `json:"mocks"`
}
