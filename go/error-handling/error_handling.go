package erratum

const testVersion = 2

// Use does things with a ResourceOpener
func Use(o ResourceOpener, input string) (err error) {
	res, err := o()
	for {
		if err == nil {
			break
		} else {
			switch err := err.(type) {
			case TransientError:
				break
			case error:
				return err
			}
		}
		res, err = o()
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				e := r.(FrobError)
				res.Defrob(e.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)
	return nil
}
