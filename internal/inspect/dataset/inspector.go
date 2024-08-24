package dataset

import (
	"bufio"
	"io"
	"os"
	"strconv"

	abytes "github.com/akramarenkov/alter/bytes"
	"github.com/akramarenkov/reusable"
	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/is"
)

// Options of inspecting. A inspected function and reader must be specified.
type Inspector[Type inspect.EightBits] struct {
	// Inspected function
	Inspected func(args ...Type) (Type, error)
	// Reader associated with dataset source
	Reader io.Reader

	// Minimum and maximum value for specified type
	min int64
	max int64

	// Buffers used to decrease allocations
	args  *reusable.Buffer[Type]
	args8 *reusable.Buffer[Type]
	items *reusable.Buffer[[]byte]

	// Result of inspecting
	result inspect.Result[Type]
}

// Validates options. A inspected function and reader must be specified.
func (insp Inspector[Type]) IsValid() error {
	if insp.Inspected == nil {
		return inspect.ErrInspectedNotSpecified
	}

	if insp.Reader == nil {
		return ErrReaderNotSpecified
	}

	return nil
}

// Performs inspecting with dataset from file.
func InspectFromFile[Type inspect.EightBits](
	path string,
	inspected func(args ...Type) (Type, error),
) (inspect.Result[Type], error) {
	file, err := os.Open(path)
	if err != nil {
		return inspect.Result[Type]{}, err
	}

	defer file.Close()

	insp := Inspector[Type]{
		Inspected: inspected,
		Reader:    file,
	}

	return insp.Inspect()
}

// Performs inspecting.
func (insp Inspector[Type]) Inspect() (inspect.Result[Type], error) {
	if err := insp.IsValid(); err != nil {
		return inspect.Result[Type]{}, err
	}

	insp.min, insp.max = inspect.PickUpRange[Type]()

	insp.args = reusable.New[Type](0)
	insp.args8 = reusable.New[Type](0)
	insp.items = reusable.New[[]byte](0)

	if err := insp.main(); err != nil {
		return inspect.Result[Type]{}, err
	}

	return insp.result, nil
}

func (insp *Inspector[Type]) main() error {
	scanner := bufio.NewScanner(insp.Reader)

	for scanner.Scan() {
		items := abytes.Split(scanner.Bytes(), []byte(" "), insp.items.Get)

		fault, reference, args, err := insp.convItems(items)
		if err != nil {
			return err
		}

		if proceed := insp.process(fault, reference, args...); !proceed {
			return scanner.Err()
		}
	}

	return scanner.Err()
}

func (insp *Inspector[Type]) convItems(items [][]byte) (bool, int64, []Type, error) {
	if len(items) <= referenceFieldsQuantity {
		return false, 0, nil, ErrNotEnoughDataInElement
	}

	fault, err := strconv.ParseBool(string(items[0]))
	if err != nil {
		return false, 0, nil, err
	}

	reference, err := strconv.ParseInt(string(items[1]), consts.DecimalBase, 64)
	if err != nil {
		return false, 0, nil, err
	}

	items = items[2:]

	args := insp.args.Get(len(items))

	for id, item := range items {
		arg, err := parseArg[Type](string(item))
		if err != nil {
			return false, 0, nil, err
		}

		args[id] = arg
	}

	return fault, reference, args, nil
}

func parseArg[Type inspect.EightBits](item string) (Type, error) {
	if is.Signed[Type]() {
		arg, err := strconv.ParseInt(item, consts.DecimalBase, 8)
		if err != nil {
			return 0, err
		}

		return Type(arg), nil
	}

	arg, err := strconv.ParseUint(item, consts.DecimalBase, 8)
	if err != nil {
		return 0, err
	}

	return Type(arg), nil
}

func (insp *Inspector[Type]) process(fault bool, reference int64, args ...Type) bool {
	// Protection against changes args from the inspected function
	copy(insp.args8.Get(len(args)), args)

	args8 := insp.args8.Get(0)

	actual, err := insp.Inspected(args8...)

	if fault {
		if err == nil {
			insp.result.Actual = actual
			insp.result.Args = args
			insp.result.Conclusion = inspect.ErrErrorExpected

			return false
		}

		insp.result.ReferenceFaults++

		return true
	}

	if reference > insp.max || reference < insp.min {
		if err == nil {
			insp.result.Actual = actual
			insp.result.Args = args
			insp.result.Conclusion = inspect.ErrErrorExpected
			insp.result.Reference = reference

			return false
		}

		insp.result.Overflows++

		return true
	}

	if err != nil {
		insp.result.Args = args
		insp.result.Conclusion = inspect.ErrUnexpectedError
		insp.result.Err = err
		insp.result.Reference = reference

		return false
	}

	if int64(actual) != reference {
		insp.result.Actual = actual
		insp.result.Args = args
		insp.result.Conclusion = inspect.ErrNotEqual
		insp.result.Reference = reference

		return false
	}

	insp.result.NoOverflows++

	return true
}
