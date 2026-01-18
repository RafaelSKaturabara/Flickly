package viewmodel

type RuleViewModel struct {
	Priority int              `json:"priority"`
	Actions  []ActionViewModel `json:"actions"`
}
