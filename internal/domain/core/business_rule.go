package core

type BusinessRule interface {
	AbleToRun() bool
	Run() bool
}
