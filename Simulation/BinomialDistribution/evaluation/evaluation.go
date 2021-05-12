package evaluation

type BinomialDitribution struct {
	Elements        int
	Probability     float64
	UnknownVariable int
}

type Distribution interface {
	NewDistribution()
}

func (binomial *BinomialDitribution) NewDistribution(elements, unknown int, probability float64) {
	binomial.Elements = elements
	binomial.UnknownVariable = unknown
	binomial.Probability = probability
}

func Factorial(value int) int {

	for i := value - 1; i >= 1; i-- {
		value *= i
	}
	return value

}

func Combinatory(n, x int) float64 {

	return float64(Factorial(n) / (Factorial(n-x) * Factorial(x)))

}
