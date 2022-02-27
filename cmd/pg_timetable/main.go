package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/cybertec-postgresql/pg_timetable/embed"
	"github.com/cybertec-postgresql/pg_timetable/embed/buildexpired"
)

func main() {
	// -----------------------------------------------------------------------------------------------------------------
	// runtime.MemProfileRate = 0
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	// -----------------------------------------------------------------------------------------------------------------
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
	// ------------------------------------------------------------------------------------------------------------
	fmt.Println("----------------------------------------------------")
	fmt.Println("---  Pg scheduler v4.4 --", buildDateTime, "---")
	fmt.Println("----------------------------------------------------")
	fmt.Println(" ")
	// ------------------------------------------------------------------------------------------------------------

	ctx, cancel := context.WithCancel(context.Background())
	embed.SetupCloseHandler(cancel)
	defer cancel()
	embed.Service(ctx)

}
