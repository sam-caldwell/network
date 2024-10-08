package network

import "fmt"

// FromString - Given a string pointer, convert it to an FQDN
func (o *Fqdn) FromString(v *string) (err error) {

	if IsValidAddress(*v) {
		*o = Fqdn(*v)
	} else {
		err = fmt.Errorf("FQDN must be valid")
	}
	return err
}
