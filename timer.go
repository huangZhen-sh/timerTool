package timerTool

import (
	"context"
	"fmt"
	"time"
)

type timerAction interface {
	Action()
}

type timerData struct {
	t      *time.Timer
	cxt    context.Context
	cancel context.CancelFunc
	tr     *time.Ticker
}

func (t timerData) Stop() {
	if t.t != nil {
		t.t.Stop()
	}
	if t.tr != nil {
		t.tr.Stop()
	}
	t.cancel()
}

func SetTimeOut(a timerAction, t time.Duration, isDebug ...bool) *timerData {
	debug := false
	if len(isDebug) > 0 {
		debug = isDebug[0]
	}
	timer := &timerData{
		t: time.NewTimer(t),
	}
	timer.cxt, timer.cancel = context.WithCancel(context.Background())
	go func(t *timerData) {
		defer func() {
			if err := recover(); err != nil {
				if debug == true {
					fmt.Println(err)
				}
			}
		}()
		select {
		case <-t.t.C:
			a.Action()
		case <-t.cxt.Done():

		}
		if debug == true {
			fmt.Printf("定时器进程已结束")
		}
	}(timer)
	return timer
}

func SetInterval(a timerAction, t time.Duration, isDebug ...bool) *timerData {
	debug := false
	if len(isDebug) > 0 {
		debug = isDebug[0]
	}
	timer := &timerData{
		tr: time.NewTicker(t),
	}
	timer.cxt, timer.cancel = context.WithCancel(context.Background())
	go func(t *timerData) {
		defer func() {
			if err := recover(); err != nil {
				if debug == true {
					fmt.Println(err)
				}
			}
		}()
		for {
			select {
			case <-t.tr.C:
				a.Action()
			case <-t.cxt.Done():
				goto End
			}
		}
	End:
		if debug == true {
			fmt.Printf("定时器进程已结束")
		}
	}(timer)
	return timer
}
