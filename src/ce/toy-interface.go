// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

type Toy interface {
	Play()

	Share()
}

type GoodToy struct {
	name string
}

type BadToy struct {
	name string
}

func (tg *GoodToy) Play() {
	fmt.Printf("%s toy is playing\n", tg.name)
}

func (tg *GoodToy) Share() {
	fmt.Printf("%s toy is sharing\n", tg.name)
}

func (tb *BadToy) Play() {
	fmt.Printf("%s toy is playing (but it does not share).\n", tb.name)
}

// We learn here that if a Type is to be abstracted out an interface, then all
// interface methods have to be necessarily implemented.
func main() {
	gt := GoodToy{"Compassionate"}

	// Note that GoodToy is accepted as Toy argument.
	play(&gt)
	share(&gt)

	// If you uncomment next 2 lines, it will give compilation error as bt
	// cannot be cast a toy since 'share' method is not implemented.
	// bt := BadToy{"Mean"}
	// play(&bt)
}

func play(t Toy) {
	t.Play()
}

func share(t Toy) {
	t.Share()
}
