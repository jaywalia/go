package org

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	uh := Org{Acr: "UH", Name: "University Of Houston"}
	fmt.Println(uh)

	rice := Org{Acr: "RICE", Name: "Rice International Center for Excellence"}
	fmt.Println(rice)

	oats := Org{Acr: "OATS", Name: "Orgnization for Arts & Technical Sciences"}
	fmt.Println(oats)

	uhCenters := uh.GetCASATestingCenters()
	if len(uhCenters) != 3 {
		t.Errorf("Found %q centers, want 3 centers", len(uhCenters))
	}

	for _, c := range uhCenters {
		switch c.Acr {
		case "GAR", "CBB", "AAH":
		default:
			t.Errorf("Unknown testing center: %q", c.Acr)
		}
	}

	oats.GetCASATestingCenters()
}
