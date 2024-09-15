package core

import "testing"

func TestAttribute_Equal(t *testing.T) {
	tests := []struct {
		lhs   Attribute
		rhs   Attribute
		equal bool
	}{
		{
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{},
			},
			equal: true,
		}, {
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x00},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{},
			},
			equal: false,
		}, {
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x00},
			},
			equal: false,
		}, {
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			equal: true,
		}, {
			lhs: Attribute{
				Type:  1,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			equal: false,
		}, {
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			rhs: Attribute{
				Type:  1,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			equal: false,
		}, {
			lhs: Attribute{
				Type:  0xFF,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			rhs: Attribute{
				Type:  0xFF,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			equal: true,
		}, {
			lhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03},
			},
			rhs: Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			equal: false,
		},
	}
	for row, tt := range tests {
		if (!tt.lhs.Equal(&tt.rhs) && tt.equal) || (tt.lhs.Equal(&tt.rhs) && !tt.equal) {
			t.Fatalf("equal method failed on %d", row)
		}
	}
}
