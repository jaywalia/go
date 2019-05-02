package org

import (
	"fmt"

	"kireikode.com/jw/center"
)

func init() {
	// fmt.Println("--------------------------------------------------------------------")
	// fmt.Println("Org package initialized. Only fired once on packet initialization.")
	// fmt.Println("--------------------------------------------------------------------")
}

// Org has acronym, name and centers
type Org struct {
	Acr  string
	Name string
}

// String functions is implementation of Stringer interface
func (o Org) String() string {
	return fmt.Sprintf("%v [%v]", o.Acr, o.Name)
}

// GetCASATestingCenters returns list of CASA Testing Centers
func (o Org) GetCASATestingCenters() []center.Center {
	switch o.Acr {
	case "UH":
		return []center.Center{
			{Acr: "GAR", Name: "CASA @ Garrison Gym"},
			{Acr: "CBB", Name: "CASA @ Classroom & Business Building"},
			{Acr: "AAH", Name: "Agnes Arnold Hall"},
		}
	default:
		return nil
	}

}
