package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Download(url string, dst io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}
	defer resp.Body.Close()

	procInfo := make(chan connInfo)

	procBar := newProcessBar(getSize(resp), procInfo)
	go procBar.start()

	if _, err := copyContent(dst, resp.Body, procInfo); err != nil {
		return err
	}
	close(procInfo)

	return nil
}

func getSize(resp *http.Response) int {
	szStr := resp.Header.Get("Content-Length")

	sz, err := strconv.Atoi(szStr)
	if err != nil {
		log.Fatalf("get size: content length not found")
	}

	return sz
}

func copyContent(dst io.Writer, src io.Reader, proccesInfo chan<- connInfo) (written int, err error) {
	buffer := make([]byte, 1000*1000)

	for {
		read, err := src.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			return written, err
		}

		w, err := dst.Write(buffer[:read])
		if err != nil {
			return written, err
		}

		written += w

		proccesInfo <- connInfo{written, time.Now()}
	}

	return written, nil
}
