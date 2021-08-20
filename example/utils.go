package main

import "time"

const (
	UTF8_CHARSET   = "UTF-8"
	SIZEOF_SHORT   = 2
	SIZEOF_INT     = 4
	SIZEOF_LONG    = 8
	NULL_BYTE      = 0x80
	TRUE_BYTE      = 1
	FALSE_BYTE     = 0
	MAX_SHIFT      = 7
	NULL_SHIFT     = 6
	BYTE_MASK      = 255
	NEXT_MASK      = -128
	NULL_NEXT_MASK = 64
	LAST_MASK      = 0
	NULL_LSB_MASK  = 63
	LSB_MASK       = 127
	TEMP_CAPACITY  = 256
)

const AgeDelta = 621355968000000

func dateFromTicks(ticks int64) time.Time {
	if ticks > 0 {

		timeT := (ticks - AgeDelta) / 10

		t := time.Unix(0, timeT*int64(time.Millisecond))

		return t

	}
	return time.Time{}
}

func dateToTicks(date time.Time) (ticks int64) {

	if !date.IsZero() {

		ticks = date.UnixNano() / int64(time.Millisecond)

		ticks = ticks*10 + AgeDelta

		return ticks

	}
	return 0
}
