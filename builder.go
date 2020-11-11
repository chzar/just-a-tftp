package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/pin/tftp"
)

func readHandler(filename string, rf io.ReaderFrom) error {

	// set the file size
	fi, err := os.Stat(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	rf.(tftp.OutgoingTransfer).SetSize(fi.Size())
	fmt.Printf("Starting Download: %s\t Size: %d\n", filename, fi.Size())

	start := time.Now()
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%s | %d bytes sent | %f Mb/s\n", filename, n, float64(n)/1000000/time.Now().Sub(start).Seconds())
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	n, err := wt.WriteTo(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes received\n", n)
	return nil
}

func BuildTftpServer(directory string, readonly bool) *tftp.Server {

	var readHandler = func(filename string, rf io.ReaderFrom) error {
		fullpath := path.Join(directory, filename)

		start := time.Now()

		file, err := os.Open(fullpath)
		if err != nil {
			logger.Errorf("%v\n", err)
			return err
		}
		n, err := rf.ReadFrom(file)
		if err != nil {
			logger.Errorf("%v\n", err)
			return err
		}
		logger.Infof("%s | %d bytes sent | %f b/s\n", filename, n, float64(n)/time.Now().Sub(start).Seconds())
		return nil
	}

	var writeHandler = func(filename string, wt io.WriterTo) error {
		fullpath := path.Join(directory, filename)

		start := time.Now()

		file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
		n, err := wt.WriteTo(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return err
		}
		logger.Infof("%s | %d bytes received | %f b/s\n", filename, n, float64(n)/time.Now().Sub(start).Seconds())
		return nil
	}

	if readonly {
		writeHandler = nil
	}

	return tftp.NewServer(readHandler, writeHandler)
}
