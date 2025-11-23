package main

type Greeter interface {
	Greet()
}

func PrintGreeting(greeters []Greeter) {
	for _, g := range greeters {
		g.Greet()
	}
}
