package main

import (
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pin/tftp"
)

func BuildTftpServer(directory string, readonly bool) *tftp.Server {

	var readHandler = func(filename string, rf io.ReaderFrom) error {
		start := time.Now()

		filename = strings.TrimLeft(filename, "/")
		fullpath := path.Join(directory, filename)

		file, err := os.Open(fullpath)
		if err != nil {
			logger.Infof("%v\n", err)
			return err
		}
		defer file.Close()

		n, err := rf.ReadFrom(file)
		// this is a mitigation for some bootloaders that interrupt their file transfer immediately after requesting it
		// if the handler is not short-circuited here and the program is running as a windows service
		// then the program will crash.
		if n == 0 {
			return nil
		}

		if err != nil {
			logger.Infof("%v\n", err)
			return err
		}
		logger.Infof("%s: %d bytes sent (%f b/s)\n", filename, n, float64(n)/time.Now().Sub(start).Seconds())
		return nil
	}

	var writeHandler = func(filename string, wt io.WriterTo) error {
		start := time.Now()

		filename = strings.TrimLeft(filename, "/")
		fullpath := path.Join(directory, filename)

		file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			logger.Infof("%v\n", err)
			return err
		}
		defer file.Close()

		n, err := wt.WriteTo(file)
		if err != nil {
			logger.Infof("%v\n", err)
			return err
		}
		logger.Infof("%s: %d bytes received (%f b/s)\n", filename, n, float64(n)/time.Now().Sub(start).Seconds())
		return nil
	}

	if readonly {
		writeHandler = nil
	}

	return tftp.NewServer(readHandler, writeHandler)
}
