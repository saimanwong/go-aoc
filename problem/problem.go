package problem

type Day string

type Year string

type Problems map[Day]Problemer

type Problemer interface {
	Run()
	Inputer
}

type Inputer interface {
	SetInput([]string)
}
