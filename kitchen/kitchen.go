package kitchen

type item int
const (
	Bread item = iota
	Cheese
	Tomato
	Lettuce

	Burger
)

var Items = []item{
	Bread,
	Cheese,
	Tomato,
	Lettuce,

	Burger,
}