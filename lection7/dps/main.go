package main

import (
	"fmt"
	"time"

	"github.com/theckman/yacspin"
)

// go get github.com/theckman/yacspin

func main() {
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[35],
		Suffix:          " backing up database to S3",
		SuffixAutoColon: true,
		Message:         "exporting data",
		StopCharacter:   "✓",
		StopColors:      []string{"fgGreen"},
	}

	spinner, err := yacspin.New(cfg)
	// handle the error

	if err != nil {
		sErr :=fmt.Errorf("some error")
		fmt.Println(sErr)
	}

	spinner.Start()

	// doing some work
	time.Sleep(2 * time.Second)

	spinner.Message("uploading data")

	// upload...
	time.Sleep(2 * time.Second)

	spinner.Stop()
}
