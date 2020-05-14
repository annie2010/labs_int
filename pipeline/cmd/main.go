package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gopherland/labs_int/pipeline/internal/pipe"
	"github.com/pkg/profile"
	prof "github.com/pkg/profile"
)

const (
	word = "gulliver"
	book = "assets/gulliver.txt"
)

func main() {
	var (
		mode      string
		profiling bool
	)

	flag.StringVar(&mode, "m", "s", "Specify mode as either s|p1|p2|p3")
	flag.BoolVar(&profiling, "p", false, "Toggles profiling")
	flag.Parse()

	if profiling {
		defer prof.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var (
		total int
		err   error
	)
	switch mode {
	case "s":
		total, err = pipe.Serial(book, word)
	case "p1":
		total, err = pipe.Pipeline(ctx, pipe.Controlled, book, word)
	case "p2":
		total, err = pipe.Pipeline(ctx, pipe.Free, book, word)
	case "p3":
		total, err = pipe.Pipeline(ctx, pipe.Buffered, book, word)
	default:
		log.Fatal("You must specify a run mode!")
	}
	if err != nil {
		log.Fatalf("Boom %v", err)
	}

	fmt.Println(total)
}
