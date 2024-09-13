package core

// FromInt - set rCtAttrCountersEnum value from integer
func (counter *CtAttrCounters) FromInt(i int) {
	*counter = CtAttrCounters(i)
}
