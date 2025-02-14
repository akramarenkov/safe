// Internal package with Incrementor which checks that the arguments change as if
// they were incremented by nested loops.
package incrementor

import (
	"errors"

	"github.com/akramarenkov/safe/internal/inspect/types"
)

var (
	ErrBeginGreaterEnd     = errors.New("begin is greater than end")
	ErrInvalidArgsQuantity = errors.New("arguments quantity exceed specified loops quantity")
	ErrInvalidSequence     = errors.New("sequence of changing argument values is broken")
)

// Checks that the arguments change as if they were incremented by nested loops.
type Incrementor[Type types.UpToI32] struct {
	begin Type
	end   Type

	expected []int64
}

// Creates Incrementor instance.
func New[Type types.UpToI32](
	loopsQuantity uint,
	begin Type,
	end Type,
) (*Incrementor[Type], error) {
	if begin > end {
		return nil, ErrBeginGreaterEnd
	}

	inc := &Incrementor[Type]{
		begin: begin,
		end:   end,

		expected: make([]int64, loopsQuantity),
	}

	inc.initExpected()

	return inc, nil
}

func (inc *Incrementor[Type]) initExpected() {
	for id := range inc.expected {
		inc.expected[id] = int64(inc.begin)
	}
}

// Performs test argument values.
func (inc *Incrementor[Type]) Test(args ...Type) error {
	if len(args) != len(inc.expected) {
		return ErrInvalidArgsQuantity
	}

	for id := range args {
		if int64(args[len(args)-id-1]) != inc.expected[id] {
			return ErrInvalidSequence
		}
	}

	inc.increase()

	return nil
}

func (inc *Incrementor[Type]) increase() {
	// smaller the id, the more nested the loop it corresponds to and vice versa
	for id := range inc.expected {
		inc.expected[id]++

		if inc.expected[id] > int64(inc.end) {
			inc.expected[id] = int64(inc.begin)
			continue
		}

		return
	}
}
