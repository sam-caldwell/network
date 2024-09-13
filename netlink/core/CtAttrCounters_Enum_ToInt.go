package core

// ToInt - return integer value fo rCtAttrCountersEnum
func (counter *CtAttrCounters) ToInt() int {
	return int(*counter)
}
