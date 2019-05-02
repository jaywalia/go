package center

import (
	"fmt"
	"time"
)

//=================================================================================
// Initialize
//=================================================================================

// load center overload hours, these overwrite default hours
func init() {
	loadTCOverloadHours()
}

//=================================================================================
// Types
//=================================================================================

// Center has a name and acronym
type Center struct {
	Acr  string
	Name string
}

// String functions is implementation of Stringer interface
func (c Center) String() string {
	return fmt.Sprintf("%v [%v]", c.Acr, c.Name)
}

// Schedule maintains open close hours
type Schedule struct {
	opens  time.Time
	closes time.Time
}

//=================================================================================

// GetTCHoursForDate returns the hours center c is open on the d date
// func GetTCHoursForDate(d time.Time, c Center) {
// 	// overload hours else Default hours
// 	//getTCOverloadHoursForDate(d, c)

// 	return
// }

// // overload hours
// func getTCOverloadHoursForDate(d time.Time, c Center) {
// 	// read from a file
// 	f, err := ioutil.ReadFile("../data/tcOverload.txt")
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	s := string(f)
// 	lines := strings.Split(s, "\n")
// 	for _, l := range lines {
// 		//fmt.Println(l)
// 		fields := strings.Split(l, ",")
// 		fmt.Println(fields)
// 		// center := fields[0]
// 		// date := fields[1]
// 		// opens := fields[2]
// 		// closes := fields[3]

// 		// fmt.Printf("%v : %v : %v : %v", center, date, opens, closes)
// 	}
// 	//fmt.Println(s)
// }

// // default schedule
// func getTCDefualtHoursForDate(d time.Time, c Center) {
// 	// default : read from a file
// }

func loadTCOverloadHours() {
	fmt.Println("load TC overload hours")

}
