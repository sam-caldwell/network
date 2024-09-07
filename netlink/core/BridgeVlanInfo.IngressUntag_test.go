package core

import "testing"

func TestIngressUntag(t *testing.T) {
	tests := []struct {
		name   string
		bridge BridgeVlanInfo
		want   bool
	}{
		{
			name:   "Untagged flag is set",
			bridge: BridgeVlanInfo{Flags: BridgeVlanInfoUntagged},
			want:   true,
		},
		{
			name:   "Untagged flag is not set",
			bridge: BridgeVlanInfo{Flags: 0},
			want:   false,
		},
		{
			name:   "Multiple flags, including untagged",
			bridge: BridgeVlanInfo{Flags: BridgeVlanInfoUntagged | 2},
			want:   true,
		},
		{
			name:   "Multiple flags, untagged not set",
			bridge: BridgeVlanInfo{Flags: 2},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bridge.IngressUntag()
			if got != tt.want {
				t.Errorf("IngressUntag() = %v, want %v", got, tt.want)
			}
		})
	}
}
