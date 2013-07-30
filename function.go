package general

type Functioner interface {
	Function() func(float64) float64
}

type Differ interface {
	Diff() func(float64) float64
}

type FuncDiffer interface {
	Functioner
	Differ
}

type Evaler interface {
	Eval(float64) (f, g float64)
}

type BasicFunc func(float64) float64

func (f BasicFunc) Function() func(float64) float64 {
	return f
}

type BasicFuncPair struct {
	F  func(float64) float64
	Fp func(float64) float64
}

func (b BasicFuncPair) Function() func(float64) float64 {
	return b.F
}

func (b BasicFuncPair) Diff() func(float64) float64 {
	return b.Fp
}
