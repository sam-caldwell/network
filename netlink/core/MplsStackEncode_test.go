package core

//
//func TestMplsStackEncode(t *testing.T) {
//	tests := []struct {
//		name     string
//		input    []MplsLabels
//		expected []byte
//	}{
//		{
//			name:     "Valid MPLS Stack Encoding",
//			input:    []MplsLabels{0x123, 0x234, 0x345, 0x456},
//			expected: []byte{0x00, 0x01, 0x23, 0x00, 0x00, 0x02, 0x34, 0x00, 0x00, 0x03, 0x45, 0x00, 0x00, 0x04, 0x56, 0x80},
//		},
//		{
//			name:     "Empty MPLS Stack",
//			input:    []MplsLabels{},
//			expected: []byte{},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := MplsStackEncode(tt.input...)
//			if len(result) != len(tt.expected) || !equalByteSlices(result, tt.expected) {
//				t.Errorf("Expected %v, got %v", tt.expected, result)
//			}
//		})
//	}
//}
//
//func equalByteSlices(a, b []byte) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	for i := range a {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}
