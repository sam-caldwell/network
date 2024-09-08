package core

// ToInt - return integer value fo rCtAttrCountersEnum
func (counter *CtAttrCountersEnum) ToInt() int {
	return int(*counter)
}
