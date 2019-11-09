package main

import (
	"fmt"
	"math"
	"strconv"
)

type Fraction struct{
	Numerator, Denominator int
}



func main() {
	var n, m, e, g int
	_, _ = fmt.Scan(&n, &m, &e, &g)
	f := Fraction{n, m}
	i := Fraction{e, g}
	if i.Denominator == 0 || f.Denominator == 0{
		fmt.Println("La calculatrice ne peut pas prendre en compte 0 comme dénominateur, merci de taper autre chose")
		main()
	}

	fmt.Println(Fraction.valeur_approchee(f))
	fmt.Println(Fraction.inverse(f))
	fmt.Println(Fraction.simplifier(f))
	fmt.Println(Fraction.diviser(f, i))
	fmt.Println(Fraction.multilplier(f, i))
	fmt.Println(Fraction.additionner(f, i))
	fmt.Println(Fraction.soustraire(f, i))
	fmt.Println(Fraction.oppose(f))
}

func (f Fraction) oppose() Fraction {
	f.Denominator = -f.Denominator
	return f
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (f Fraction) simplifier() Fraction{
	gcd := GCD(f.Numerator, f.Denominator)
	f.Numerator /= gcd
	f.Denominator /= gcd

	return f
}

func (f Fraction) valeur_approchee() float64{
	y := float64(f.Numerator) / float64(f.Denominator)
	return math.Round(y)
}

func (f Fraction) soustraire(g Fraction) Fraction{
	return f.additionner(Fraction{-g.Numerator, g.Denominator})
}

func (f Fraction) inverse() Fraction{
	return Fraction{f.Denominator, f.Numerator}
}

func (f Fraction) diviser(g Fraction) Fraction{
	return f.multilplier(Fraction{g.Denominator, g.Numerator})
}

func (f Fraction) multilplier(g Fraction) Fraction{
	f.Numerator *= g.Numerator
	f.Denominator *= g.Denominator
	return f
}

func (f Fraction) additionner(b Fraction) Fraction {
		num := f.Numerator*b.Denominator + b.Numerator*f.Denominator
		den := f.Denominator * b.Denominator
		d := GCD(abs(num), den)
		return Fraction{
		Numerator: num / d,
		Denominator: den / d,
	}
}

func (f Fraction) String() string{
		return strconv.Itoa(f.Numerator) + "/" + strconv.Itoa(f.Denominator)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
