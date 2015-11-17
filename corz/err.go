// Shows how to deal with errors, panic attacks and runtime exceptions
//
// Author: Dmitri Krasnenko

package corz

import "errors"

func Div(i,c int64) float64 {
	//Will fail on runtime exception when c = 0
	return float64(i/c)
}

func DivPanic(i,c int64) (float64) {
	//Fall with panic.
	if c == 0 {
		panic("Error: integer divide by zero.")
	}

	return float64(i/c)
}

func DivErr(i,c int64) (float64, error)  {
	//Naive way to treat division by zero case
	if c == 0 {
		return 0, errors.New("Error...")
	}

	return float64(i/c), nil
}

func DivSafe(i,c int64) (res float64, r error) {
	//More sophisticated way. Catches any runtime errors.
	defer func() {
		if err := recover(); err != nil {
			r = err.(error)
            res = 0
        }
	}()

	return float64(i/c), nil
}


func NewEmployeeSafe(id int64, name, position string) (*Employee, error) {
	if name == "" || len(position) == 0 {
		return nil, errors.New("Invalid input parameters.")
	}

	return &Employee{ NewPerson(id, name),
		position,
	}, nil
}
