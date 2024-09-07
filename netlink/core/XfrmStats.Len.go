package core

// Len - return size of XfrmStats struct (in bytes)
func (msg *XfrmStats) Len() int {
	return SizeOfXfrmStats
}
