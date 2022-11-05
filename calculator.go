package calculator

type Calculator struct {
	Result int
}

func New() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(num int) *Calculator {
	c.Result += num
	return c
}

func (c *Calculator) Sub(num int) *Calculator {
	c.Result -= num
	return c
}

func (c *Calculator) Mul(num int) *Calculator {
	c.Result *= num
	return c
}

func (c *Calculator) Div(num int) *Calculator {
	c.Result /= num
	return c
}

func (c *Calculator) Calculate() int {
	return c.Result
}
