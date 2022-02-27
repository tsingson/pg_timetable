package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cybertec-postgresql/pg_timetable/embed/buildexpired"
)

func main() {
	buildDateTime, err := buildexpired.BuildDateTime()
	if err != nil {
		os.Exit(2)
	}
	expired := buildDateTime.Add(buildexpired.Expired)

	t1 := time.Now().UTC()
	if !t1.Before(expired) {
		fmt.Println("Expired at ", expired)
		os.Exit(2)
	}
	fmt.Println("good luck at ", expired)
}
