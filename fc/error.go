package fc

import "errors"

func ErrSlice2Err(es []error) error {
	estr := ""
	for _, e := range es {
		if e != nil {
			estr += e.Error() + "\n"
		}
	}
	if estr == "" {
		return nil
	}
	return errors.New(estr)
}
