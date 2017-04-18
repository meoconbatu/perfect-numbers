package perfect

import "errors"

const testVersion = 1

type Classification int

const (
	ClassificationPerfect Classification = iota + 1
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Error")

func Classify(input uint64) (output Classification, err error) {
	if input < 1 {
		err = ErrOnlyPositive
		return
	}
	if input == 1 {
		output = ClassificationDeficient
		return
	}
	sum := uint64(1)
	for i, t := uint64(2), input; i < t; i++ {
		if input%i == 0 {
			sum += i + input/i
			if sum > input {
				output = ClassificationAbundant
				return
			}
			t = input / i

		}
	}
	if input == sum {
		output = ClassificationPerfect
	} else if input > sum {
		output = ClassificationDeficient
	} else {
		output = ClassificationAbundant
	}
	return
}
