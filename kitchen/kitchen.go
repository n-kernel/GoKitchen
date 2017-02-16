package kitchen

type Item int
const (
	Bread   Item = iota
	Cheese
	Tomato
	Lettuce

	Burger
)

var Items = []Item{
	Bread,
	Cheese,
	Tomato,
	Lettuce,

	Burger,
}