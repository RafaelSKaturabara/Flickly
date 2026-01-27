package viewmodel

type ConditionViewModel struct {
	Priority        int    `json:"priority"`
	FieldToCompare  string `json:"fieldToCompare"`
	Operator        string `json:"operator"`
	ValueToCompare  string `json:"valueToCompare"`
}
