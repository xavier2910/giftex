package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	filename := flag.String("input", "people.json", "A json file containing the people in the gift exchange")
	flag.Parse()

	fmt.Printf("opening %s...", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("ERROR")
		return err
	}

	people, err := decodeJsonFile(f)
	if err != nil {
		fmt.Println("ERROR")
		return err
	}
	fmt.Println("OK")

	fmt.Print("authenticating...")
	auth, sender, err := authorize()
	if err != nil {
		fmt.Println("ERROR")
		return err
	}
	fmt.Println("OK")

	gmp := genNames(people)

	var eg errgroup.Group

	for giver, recipient := range gmp {
		g := giver
		r := recipient

		eg.Go(func() error {
			fmt.Printf("emailing %s... \n", g.Name)
			err := email(r, g, auth, *sender)
			if err != nil {
				fmt.Printf("ERROR emailing %s. see end of output for details", g.Name)
				return err
			} else {
				fmt.Printf("OK, emailed %s", g.Name)
				return nil
			}
		})
	}

	return eg.Wait()
}

func genNames(people []person) (givingMap map[person]person) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	givingMap = make(map[person]person, 4)

	rindices := r.Perm(len(people))

	for i := 0; i < len(people); i++ {

		var prev int
		if i == 0 {
			prev = len(people) - 1
		} else {
			prev = i - 1
		}

		givingMap[people[rindices[prev]]] = people[rindices[i]]

	}

	return
}
