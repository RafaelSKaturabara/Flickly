package entities

type Condition struct {
	Prioridade        int
	FieldToCompare    string
	Operator          OperatorToCompare
	ValueToCompare    string
}
