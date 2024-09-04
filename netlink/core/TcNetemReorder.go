package core

// TcNetemReorder - tc_netem_reorder - Structure for netem reordering options.
type TcNetemReorder struct {

	// Probability - reordering probability (0=none ~0=100%)
	Probability uint32

	// Correlation - reordering correlation
	Correlation uint32
}
