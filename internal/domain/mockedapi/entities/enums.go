package entities

type HttpMethod string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodHead    HttpMethod = "HEAD"
)

type OperatorToCompare string

const (
	OperatorEqual        OperatorToCompare = "=="
	OperatorNotEqual     OperatorToCompare = "!="
	OperatorGreater      OperatorToCompare = ">"
	OperatorGreaterEqual OperatorToCompare = ">="
	OperatorLess         OperatorToCompare = "<"
	OperatorLessEqual    OperatorToCompare = "<="
	OperatorContains     OperatorToCompare = "contains"
	OperatorHasSuffix    OperatorToCompare = "hasSuffix"
	OperatorHasPrefix    OperatorToCompare = "hasPrefix"
)
