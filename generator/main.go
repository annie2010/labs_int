package main

import (
	"fmt"
	"log"

	"github.com/gookit/color"
	"github.com/gopherland/labs_int/generator/stacks"
)

//go:generate stacker -t float64,int32
func main() {
	fmt.Print("\033[H\033[2J")

	tryFloat()
	tryInt()
}

func tryFloat() {
	s := stacks.Float64{}
	for _, v := range []float64{10.5, 20.2, 42.25} {
		s.Push(v)
	}
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	log.Printf("🥞 <<%5s>> Pop:%v -- Top:%v -- Peek:%v", "Float", cyan(v), green(s.Top()), &s)
}

func tryInt() {
	s := stacks.Int32{}
	for _, v := range []int32{200, 100, 300} {
		s.Push(v)
	}
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	log.Printf("📚 <<%5s>> Pop:%v -- Top:%v -- Peek:%v", "Float", cyan(v), green(s.Top()), &s)
}
