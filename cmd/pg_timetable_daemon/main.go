package main

import (
	"context"
	"fmt"
	"github.com/cybertec-postgresql/pg_timetable/embed"
	"github.com/cybertec-postgresql/pg_timetable/embed/buildexpired"
	"github.com/cybertec-postgresql/pg_timetable/embed/load"
	"github.com/sevlyar/go-daemon"
	"os"
	"runtime"
	"time"
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

	// ------------------------------------------------------------------------------------------------------------

	path, _ := load.GetCurrentExecDir()

	stopSignal := make(chan struct{})

	// undo := zap.RedirectStdLog(ws.Log.Log)
	// defer undo()

	me := os.Args[0]
	pidFile := me + ".pid"
	// ------------------------------------------------------------------------------------------------------------
	ctxt := &daemon.Context{
		PidFileName: pidFile,
		PidFilePerm: 0o644,
		// LogFileName: path + "/log/" + logFile,
		LogFilePerm: 0640,
		WorkDir:     path,
		Umask:       0o27,
		Args:        []string{"msa-v3"},
	}
	d, er1 := ctxt.Reborn()
	if er1 != nil {
		fmt.Printf("Error %v", er1)
	}
	if d != nil {
		return
	}
	// nolint
	defer func(ctxt *daemon.Context) {
		err := ctxt.Release()
		if err != nil {
			os.Exit(2)
		}
	}(ctxt)
	// -----------------------------------------------------------------------------------------------------------------
	embed.Service(ctx)
	// select {}
	<-stopSignal
}
