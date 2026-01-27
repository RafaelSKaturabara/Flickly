package viewmodel

type HttpRequestViewModel struct {
	Method      string               `json:"method"`
	Path        string               `json:"path"`
	Headers     []HeaderViewModel    `json:"headers"`
	QueryParams []QueryParamViewModel `json:"queryParams"`
	Body        *BodyViewModel       `json:"body"`
	Route       *RouteViewModel      `json:"route"`
}
