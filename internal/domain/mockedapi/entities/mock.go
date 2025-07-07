package entities

type Mock struct {
	Conditions []Condition
	Rules      []Rule
	Response   *HttpResponse
	HasCondition bool
}
