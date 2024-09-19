package bell

import (
	"context"
	"testing"
	"time"
)

const sendDefaultSize = 200
const sendDefaultTime = time.Second * 10

func TestAlarmBell(t *testing.T) {
	bell := NewBell(sendDefaultTime, sendDefaultSize)
	go bell.start(context.Background())
	for i := 0; i < sendDefaultSize-1; i++ {
		if bell.needAlarm() {
			t.Error()
		}
	}
}

func TestAlarmBell2(t *testing.T) {
	bell := NewBell(sendDefaultTime, sendDefaultSize)
	go bell.start(context.Background())

	count := 0
	for i := 0; i < sendDefaultSize*3; i++ {
		if bell.needAlarm() {
			count++
		}
	}
	if count != 1 {
		t.Error()
	}
}

func TestAlarmBell3(t *testing.T) {
	ts := time.Millisecond * 200
	bell := NewBell(ts, sendDefaultSize)
	go bell.start(context.Background())

	count := 0
	for i := 0; i < sendDefaultSize*3; i++ {
		if bell.needAlarm() {
			count++
			time.Sleep(ts)
		}
	}
	if count <= 1 {
		t.Error(count)
	}
}
