package core

// Tcf represents timing data related to a Traffic Control Filter.
//
// See https://tldp.org/HOWTO/Traffic-Control-HOWTO/components.html,
// https://events.static.linuxfound.org/sites/events/files/slides/Linux_traffic_control.pdf
type Tcf struct {
	Install  uint64 // Install timestamp or the time the filter was installed.
	LastUse  uint64 // LastUse represents the last time the filter was used.
	Expires  uint64 // Expires indicates when the filter will expire.
	FirstUse uint64 // FirstUse is the timestamp of the first use of the filter.
}
