package main

import (
	"fmt"

	"kireikode.com/jw/org"
)

func main() {

	//sd := time.Date(2019, 04, 21, 0, 0, 0, 0, time.UTC)
	//ed := time.Date(2019, 04, 28, 0, 0, 0, 0, time.UTC)

	// fmt.Println("scheduler!")
	// fmt.Println(string.Reverse("hello!"))
	// for _, c := range center.GetCASATestingCenters() {
	// 	fmt.Println(c)
	// }
	org := org.Org{Acr: "UH", Name: "University Of Houston"}
	centers := org.GetCASATestingCenters()
	for _, c := range centers {
		fmt.Println(c)
		//center.GetTCHoursForDate(sd, c)
	}

	// from start date to end date
	// for d := sd; d.Before(ed); d = d.AddDate(0, 0, 1) {
	// 	// get testing centers
	// 	{
	// 		// check to see what are the hours for the testing center
	// 		// for _, s := range center.GetTCHoursForDate(d, c) {
	// 		// 	fmt.Println("On %v, day: %v, %s is open from %s to %s", d, d.Weekday(), c, s.opens, s.closes)
	// 		// }
	// 	}

	// 	// now we pick up workers
	// 	// 		need to loop through workers
	// 	//		break time slots into 15 minutes
	// 	//
	// 	fmt.Println(d)
	// }
}
