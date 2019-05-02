package datamgr

import (
	"fmt"
	"io/ioutil"
)

func init() {
	fmt.Println("data init ...")
}

// ReadSpecialSchedule reads the special schedule file
func ReadSpecialSchedule(p string) (string, error) {
	fc, e := ioutil.ReadFile(p)
	if e != nil {
		panic(e)
		// maybe log the error
	}
	return string(fc), nil
}
