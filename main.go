//go:build !lambda
// +build !lambda

package main

func main() {
	r := Entrypoint(false)

	r.Run("localhost:4000")
}
