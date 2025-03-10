package safe

import (
	"slices"

	"github.com/akramarenkov/safe/internal/clone"
	"github.com/akramarenkov/safe/internal/is"

	"golang.org/x/exp/constraints"
)

var pow10Table = [...]uint64{ //nolint:gochecknoglobals // To increase the performance
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
}

// Adds three integers and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Add3[Type constraints.Integer](first, second, third Type) (Type, error) {
	if interim, err := Add(first, second); err == nil {
		if sum, err := Add(interim, third); err == nil {
			return sum, nil
		}
	}

	if interim, err := Add(first, third); err == nil {
		if sum, err := Add(interim, second); err == nil {
			return sum, nil
		}
	}

	if interim, err := Add(second, third); err == nil {
		return Add(first, interim)
	}

	return 0, ErrOverflow
}

// Adds three unsigned integers and detects whether an overflow has occurred or not.
//
// Faster than the [Add3] function about 70%.
//
// In case of overflow, an error is returned.
func Add3U[Type constraints.Unsigned](first, second, third Type) (Type, error) {
	interim, err := AddU(first, second)
	if err != nil {
		return 0, err
	}

	return AddU(interim, third)
}

// Adds up several integers and detects whether an overflow has occurred or not.
//
// Slower than the [Add] function about 220%, than the [Add3] function about 65%. And
// overall very slow, be careful.
//
// In case of overflow or missing arguments, an error is returned.
func AddM[Type constraints.Integer](addends ...Type) (Type, error) {
	//nolint:mnd // Adding constants will not improve readability for checking the
	// number of arguments
	switch len(addends) {
	case 0:
		return 0, ErrMissingArguments
	case 1:
		return addends[0], nil
	case 2:
		return Add(addends[0], addends[1])
	case 3:
		return Add3(addends[0], addends[1], addends[2])
	}

	addends = clone.Slice(addends)

	sorted := false

	for len(addends) != 3 {
		interim, err := Add(addends[0], addends[len(addends)-1])
		if err != nil {
			if !sorted {
				sorted = true

				slices.Sort(addends)

				continue
			}

			return 0, err
		}

		sorted = false

		addends[0] = interim
		addends = addends[:len(addends)-1]
	}

	return Add3(addends[0], addends[1], addends[2])
}

// Adds up several unsigned integers and detects whether an overflow has occurred or
// not.
//
// Slower than the [AddU] function about 170%, than the [Add3U] function about 80%,
// faster than the [AddM] function about 85%.
//
// In case of overflow or missing arguments, an error is returned.
func AddMU[Type constraints.Unsigned](addends ...Type) (Type, error) {
	if len(addends) == 0 {
		return 0, ErrMissingArguments
	}

	sum := addends[0]

	for _, addend := range addends[1:] {
		interim, err := AddU(sum, addend)
		if err != nil {
			return 0, err
		}

		sum = interim
	}

	return sum, nil
}

// Subtracts three integers (subtrahend, deductible from minuend) and detects whether
// an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Sub3[Type constraints.Integer](minuend, subtrahend, deductible Type) (Type, error) {
	if interim, err := Sub(minuend, subtrahend); err == nil {
		diff, err := Sub(interim, deductible)
		if err == nil {
			return diff, nil
		}
	}

	if interim, err := Sub(minuend, deductible); err == nil {
		if diff, err := Sub(interim, subtrahend); err == nil {
			return diff, nil
		}
	}

	if interim, err := Add(subtrahend, deductible); err == nil {
		return Sub(minuend, interim)
	}

	return 0, ErrOverflow
}

// Subtracts three unsigned integers (subtrahend, deductible from minuend) and
// detects whether an overflow has occurred or not.
//
// Faster than the [Sub3] function about 75%.
//
// In case of overflow, an error is returned.
func Sub3U[Type constraints.Unsigned](minuend, subtrahend, deductible Type) (Type, error) {
	interim, err := SubU(minuend, subtrahend)
	if err != nil {
		return 0, err
	}

	return SubU(interim, deductible)
}

// Subtracts several integers (subtrahends from minuend) and detects whether an
// overflow has occurred or not.
//
// Slower than the [Sub] function about 160%, than the [Sub3] function about 40%. And
// overall very slow, be careful.
//
// In case of overflow, an error is returned.
func SubM[Type constraints.Integer](minuend Type, subtrahends ...Type) (Type, error) {
	//nolint:mnd // Adding constants will not improve readability for checking the
	// number of arguments
	switch len(subtrahends) {
	case 0:
		return minuend, nil
	case 1:
		return Sub(minuend, subtrahends[0])
	case 2:
		return Sub3(minuend, subtrahends[0], subtrahends[1])
	}

	subtrahends = clone.Slice(subtrahends)

	for len(subtrahends) != 2 {
		found := false
		maximum := Type(0)
		maximumID := 0

		for id, subtrahend := range subtrahends {
			interim, err := Sub(minuend, subtrahend)
			if err != nil {
				continue
			}

			if interim > maximum || !found {
				found = true
				maximum = interim
				maximumID = id
			}
		}

		if found {
			minuend = maximum
			subtrahends[0], subtrahends[maximumID] = subtrahends[maximumID], subtrahends[0]
			subtrahends = subtrahends[1:]

			continue
		}

		return 0, ErrOverflow
	}

	return Sub3(minuend, subtrahends[0], subtrahends[1])
}

// Subtracts several unsigned integers (subtrahends from minuend) and detects
// whether an overflow has occurred or not.
//
// Slower than the [SubU] function about 90%, than the [Sub3U] function about 100%,
// faster than the [SubM] function about 85%.
//
// In case of overflow, an error is returned.
func SubMU[Type constraints.Unsigned](minuend Type, subtrahends ...Type) (Type, error) {
	diff := minuend

	for _, subtrahend := range subtrahends {
		interim, err := SubU(diff, subtrahend)
		if err != nil {
			return 0, err
		}

		diff = interim
	}

	return diff, nil
}

// Multiplies three integers and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Mul3[Type constraints.Integer](first, second, third Type) (Type, error) {
	if interim, err := Mul(first, second); err == nil {
		if product, err := Mul(interim, third); err == nil {
			return product, nil
		}
	}

	if interim, err := Mul(first, third); err == nil {
		if product, err := Mul(interim, second); err == nil {
			return product, nil
		}
	}

	if interim, err := Mul(second, third); err == nil {
		return Mul(first, interim)
	}

	return 0, ErrOverflow
}

// Multiplies three unsigned integers and detects whether an overflow has occurred or
// not.
//
// Faster than the [Mul3] function about 65%.
//
// In case of overflow, an error is returned.
func Mul3U[Type constraints.Unsigned](first, second, third Type) (Type, error) {
	if third == 0 {
		return 0, nil
	}

	interim, err := MulU(first, second)
	if err != nil {
		return 0, err
	}

	return MulU(interim, third)
}

// Multiplies several integers and detects whether an overflow has occurred or not.
//
// Slower than the [Mul] function about 10%. And overall very slow, be careful.
//
// In case of overflow or missing arguments, an error is returned.
func MulM[Type constraints.Integer](factors ...Type) (Type, error) {
	//nolint:mnd // Adding constants will not improve readability for checking the
	// number of arguments
	switch len(factors) {
	case 0:
		return 0, ErrMissingArguments
	case 1:
		return factors[0], nil
	case 2:
		return Mul(factors[0], factors[1])
	case 3:
		return Mul3(factors[0], factors[1], factors[2])
	}

	factors = clone.Slice(factors)

	sortMulM(factors)

	for _, factor := range factors {
		if factor == 0 {
			return 0, nil
		}
	}

	product := factors[0]

	for _, factor := range factors[1:] {
		interim, err := Mul(product, factor)
		if err != nil {
			return 0, err
		}

		product = interim
	}

	return product, nil
}

// Multiplies several unsigned integers and detects whether an overflow has
// occurred or not.
//
// Slower than the [Mul3U] function about 40%, faster than the [MulM] function about 70%.
//
// In case of overflow or missing arguments, an error is returned.
func MulMU[Type constraints.Unsigned](factors ...Type) (Type, error) {
	if len(factors) == 0 {
		return 0, ErrMissingArguments
	}

	for _, factor := range factors {
		if factor == 0 {
			return 0, nil
		}
	}

	product := factors[0]
	factors = factors[1:]

	for _, factor := range factors {
		interim, err := MulU(product, factor)
		if err != nil {
			return 0, err
		}

		product = interim
	}

	return product, nil
}

// Divides several integers (dividend to divisors) and detects whether an overflow
// has occurred or not.
//
// The divisors is also checked for equality to zero.
//
// Slower than the [Div] function about 15%.
//
// In case of overflow or divisors equal to zero, an error is returned.
func DivM[Type constraints.Integer](dividend Type, divisors ...Type) (Type, error) {
	quotient := dividend

	minusOnes := 0

	for _, divisor := range divisors {
		if divisor == 0 {
			return 0, ErrDivisionByZero
		}

		if is.MinusOne(divisor) {
			minusOnes++
			continue
		}

		quotient /= divisor
	}

	// Paired minus ones cancel each other out
	if is.Even(minusOnes) {
		return quotient, nil
	}

	if is.Min(quotient) {
		return 0, ErrOverflow
	}

	return -quotient, nil
}

// Raises 10 to a power and detects whether an overflow has occurred or not.
//
// In case of overflow, an error is returned.
func Pow10[Type, TypePower constraints.Integer](power TypePower) (Type, error) {
	if power < 0 {
		return 0, nil
	}

	// Value of pow10table length fits into any integer type
	if power >= TypePower(len(pow10Table)) {
		return 0, ErrOverflow
	}

	return IToI[Type](pow10Table[power])
}

// Raises base to a power and detects whether an overflow has occurred or not.
//
// Straightforward and slow implementation. Be careful.
//
// In case of overflow, an error is returned.
func Pow[Type, TypePower constraints.Integer](base Type, power TypePower) (Type, error) {
	if power == 0 {
		return 1, nil
	}

	if base == 1 {
		return 1, nil
	}

	if base == 0 {
		if power < 0 {
			return 0, ErrDivisionByZero
		}

		return 0, nil
	}

	if power < 0 {
		if is.MinusOne(base) {
			if is.Even(power) {
				return 1, nil
			}

			return base, nil
		}

		return 0, nil
	}

	powered := base

	for step := TypePower(1); step < power; step++ {
		// Overflow must be checked at each multiplication step
		product, err := Mul(powered, base)
		if err != nil {
			return 0, err
		}

		powered = product
	}

	return powered, nil
}
