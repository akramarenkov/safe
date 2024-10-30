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
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/internal/is"
	"github.com/akramarenkov/safe/intspec"
)

// Options of inspecting. A inspected function and reader must be specified.
type Inspector[Type types.UpToUSI32] struct {
	// Inspected function
	Inspected func(args ...Type) (Type, error)
	// Reader associated with dataset source
	Reader io.Reader

	// Minimum and maximum value for specified type
	minimum int64
	maximum int64

	// Buffers used to decrease allocations
	args    *reusable.Buffer[Type]
	argsDup *reusable.Buffer[Type]
	fields  *reusable.Buffer[[]byte]

	// Result of inspecting
	result types.Result[Type, Type, int64]
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
func InspectFromFile[Type types.UpToUSI32](
	path string,
	inspected func(args ...Type) (Type, error),
) (types.Result[Type, Type, int64], error) {
	file, err := os.Open(path)
	if err != nil {
		return types.Result[Type, Type, int64]{}, err
	}

	defer file.Close()

	insp := Inspector[Type]{
		Inspected: inspected,
		Reader:    file,
	}

	return insp.Inspect()
}

// Performs inspecting.
func (insp Inspector[Type]) Inspect() (types.Result[Type, Type, int64], error) {
	if err := insp.IsValid(); err != nil {
		return types.Result[Type, Type, int64]{}, err
	}

	insp.minimum, insp.maximum = inspect.ConvSpan[Type, int64]()

	insp.args = reusable.New[Type](0)
	insp.argsDup = reusable.New[Type](0)
	insp.fields = reusable.New[[]byte](0)

	if err := insp.main(); err != nil {
		return types.Result[Type, Type, int64]{}, err
	}

	return insp.result, nil
}

func (insp *Inspector[Type]) main() error {
	scanner := bufio.NewScanner(insp.Reader)

	for scanner.Scan() {
		fields := abytes.Split(scanner.Bytes(), []byte(" "), insp.fields.Get)

		fault, reference, args, err := insp.convFields(fields)
		if err != nil {
			return err
		}

		if proceed := insp.process(fault, reference, args...); !proceed {
			return scanner.Err()
		}
	}

	return scanner.Err()
}

func (insp *Inspector[Type]) convFields(fields [][]byte) (bool, int64, []Type, error) {
	if len(fields) <= referenceFieldsQuantity {
		return false, 0, nil, ErrNotEnoughDataInItem
	}

	fault, err := strconv.ParseBool(string(fields[0]))
	if err != nil {
		return false, 0, nil, err
	}

	reference, err := strconv.ParseInt(string(fields[1]), consts.DecimalBase, 64)
	if err != nil {
		return false, 0, nil, err
	}

	fields = fields[2:]

	args := insp.args.Get(len(fields))

	for id, field := range fields {
		arg, err := parseArg[Type](string(field))
		if err != nil {
			return false, 0, nil, err
		}

		args[id] = arg
	}

	return fault, reference, args, nil
}

func parseArg[Type types.UpToUSI32](field string) (Type, error) {
	size := intspec.BitSize[Type]()

	if is.Signed[Type]() {
		arg, err := strconv.ParseInt(field, consts.DecimalBase, size)
		if err != nil {
			return 0, err
		}

		return Type(arg), nil
	}

	arg, err := strconv.ParseUint(field, consts.DecimalBase, size)
	if err != nil {
		return 0, err
	}

	return Type(arg), nil
}

func (insp *Inspector[Type]) process(fault bool, reference int64, args ...Type) bool {
	// Protection against changes args from the inspected function
	copy(insp.argsDup.Get(len(args)), args)

	actual, err := insp.Inspected(insp.argsDup.Get(0)...)

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

	if reference > insp.maximum || reference < insp.minimum {
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
