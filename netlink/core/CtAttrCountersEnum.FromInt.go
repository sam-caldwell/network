package core

// FromInt - set rCtAttrCountersEnum value from integer
func (counter *CtAttrCountersEnum) FromInt(i int) {
	*counter = CtAttrCountersEnum(i)
}
