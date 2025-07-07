package viewmodel

type MockViewModel struct {
	Conditions   []ConditionViewModel `json:"conditions"`
	Rules        []RuleViewModel      `json:"rules"`
	Response     *HttpResponseViewModel `json:"response"`
	HasCondition bool                 `json:"hasCondition"`
}

