package combine

import "github.com/hashicorp/go-multierror"

func Errors(errs ...error) error {
	if errs ==nil{
		return nil
	}
	if len(errs) == 0 {
		return nil
	}
	var err error = nil
	for _, e := range errs {
		if e == nil {
			continue
		}
		if err == nil {
			err = e
		} else {
			err = multierror.Append(err, e)
		}
	}

	return err
}
