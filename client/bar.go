package client

import (
	"fmt"
	"time"
)

func startProccessBar(size int, procInfo <-chan connInfo) {
	start := time.Now()
	tick := time.Tick(time.Millisecond * 400)

	for {
		select {
		case <-tick:
			cur, ok := <-procInfo
			if !ok {
				return
			}

			speed := getSpeed(start, cur)
			fmt.Printf("                                                                     \r")
			fmt.Printf("[%v]\t%d %.1f Kb/s %.1f s.\r", getProccessString(size, cur.written),
				getProccessInPercent(size, cur.written), speed, getRemainingTime(speed, size-cur.written))
		}
	}
}

func getProccessString(size, cur int) string {
	const percentInOneChar = 5

	procstr := make([]byte, int(100/percentInOneChar))
	percentOfDownloaded := getProccessInPercent(size, cur)

	for i, _ := range procstr {
		if (i+1)*percentInOneChar < int(percentOfDownloaded) {
			procstr[i] = '#'
		} else {
			procstr[i] = ' '
		}
	}

	return string(procstr)
}

func getSpeed(start time.Time, cur connInfo) float64 {
	return float64(cur.written/1000) / float64(cur.t.Sub(start)/time.Second)
}

func getRemainingTime(speed float64, remainingSize int) float64 {
	return float64(remainingSize/1000) / speed
}

func getProccessInPercent(size, cur int) int {
	return int(float64(cur) / float64(size) * 100)
}
