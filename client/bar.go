package client

import (
	"fmt"
	"time"
)

type processBar struct {
	size        int
	startOfTime time.Time
	cur         connInfo
	procInfo    <-chan connInfo
	ticker      *time.Ticker
}

func newProcessBar(size int, procInfo <-chan connInfo) *processBar {
	return &processBar{
		size:        size,
		startOfTime: time.Now(),
		procInfo:    procInfo,
		ticker:      time.NewTicker(time.Millisecond * 400),
	}
}

func (p *processBar) start() {
	defer p.ticker.Stop()
	var ok bool

	for {
		select {
		case <-p.ticker.C:
			p.printInfo()

		case p.cur, ok = <-p.procInfo:
			if !ok {
				return
			}
		}
	}
}

func (p *processBar) printInfo() {
	proc := p.getProccessInPercent()
	procStr := newProcessingString(5)

	fmt.Printf("                                                                     \r")
	fmt.Printf("[%v]\t%d%% %.1f Kb/s %.1f s.\r", procStr.get(proc), proc, p.getSpeed(), p.getRemainingTime())
}

func (p *processBar) getProccessInPercent() int {
	return int(float64(p.cur.written) / float64(p.size) * 100)
}

func (p *processBar) getSpeed() float64 {
	return float64(p.cur.written/1000) / float64(p.cur.t.Sub(p.startOfTime)/time.Second)
}

func (p *processBar) getRemainingTime() float64 {
	return float64((p.size-p.cur.written)/1000) / p.getSpeed()
}
