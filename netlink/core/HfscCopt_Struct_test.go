package core

import (
	"testing"
)

// TestHfscCopt tests the HfscCopt struct and the functionality of the Curve methods.
func TestHfscCopt(t *testing.T) {
	// Initialize an HfscCopt instance with default values
	hfsc := HfscCopt{
		Rsc: Curve{m1: 1000, d: 500, m2: 2000},
		Fsc: Curve{m1: 3000, d: 1500, m2: 4000},
		Usc: Curve{m1: 5000, d: 2500, m2: 6000},
	}

	// Check if the initial values of Rsc, Fsc, and Usc are correct using the Attrs method
	checkCurveAttrs(t, hfsc.Rsc, 1000, 500, 2000, "Rsc")
	checkCurveAttrs(t, hfsc.Fsc, 3000, 1500, 4000, "Fsc")
	checkCurveAttrs(t, hfsc.Usc, 5000, 2500, 6000, "Usc")

	// Update the Rsc, Fsc, and Usc curves using the Set method
	hfsc.Rsc.Set(1100, 600, 2100)
	hfsc.Fsc.Set(3100, 1600, 4100)
	hfsc.Usc.Set(5100, 2600, 6100)

	// Check if the updated values of Rsc, Fsc, and Usc are correct
	checkCurveAttrs(t, hfsc.Rsc, 1100, 600, 2100, "Rsc (updated)")
	checkCurveAttrs(t, hfsc.Fsc, 3100, 1600, 4100, "Fsc (updated)")
	checkCurveAttrs(t, hfsc.Usc, 5100, 2600, 6100, "Usc (updated)")
}

// checkCurveAttrs is a helper function to check the attributes of a Curve.
func checkCurveAttrs(t *testing.T, curve Curve, expectedM1, expectedD, expectedM2 uint32, curveName string) {
	m1, d, m2 := curve.Attrs()
	if m1 != expectedM1 || d != expectedD || m2 != expectedM2 {
		t.Errorf("For %s, expected (m1: %d, d: %d, m2: %d) but got (m1: %d, d: %d, m2: %d)",
			curveName, expectedM1, expectedD, expectedM2, m1, d, m2)
	}
}
