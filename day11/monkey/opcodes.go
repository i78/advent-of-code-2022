package monkey

type Operation interface {
	Exec(int) int
}

type Add struct {
	Operand int
}

func (a Add) Exec(n int) int {
	return a.Operand + n
}

type Mul struct {
	Operand int
}

func (m Mul) Exec(n int) int {
	return m.Operand * n
}

type Pow2 struct {
	Operand int
}

func (d Pow2) Exec(n int) int {
	return n*n
}