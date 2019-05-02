package datamgr

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("data test....")

	d, e := os.Getwd()
	panicOnErr(e)
	fmt.Println(d)

	fc, err := ReadSpecialSchedule("tcSpecialHours.txt")
	panicOnErr(err)
	//fmt.Println(string(fc))

	lines := strings.Split(fc, "\n")
	for _, l := range lines {
		fmt.Println(l)
		cols := strings.Split(l, ",")
		for _, c := range cols {
			fmt.Println(c)
		}
		fmt.Println("-------------")
		break
	}
}

func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}
