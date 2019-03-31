// 048 epoch.go

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	now := time.Now()
	secs := now.Unix()

	nanos := now.UnixNano()
	const _1M = 1000 * 1000
	millis := nanos / _1M

	p(now)
	p(secs)
	p(millis)
	p(nanos)

	p(time.Unix(secs, 0))
	p(time.Unix(0, nanos))
}
