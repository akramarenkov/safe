package dataset

import (
	"io"
	"os"
	"strconv"

	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset/filler"
)

// Options of collecting. A reference function and writer must be specified.
type Collector[Type inspect.EightBits] struct {
	// Quantity of arguments for inspected and reference functions
	ArgsQuantity int
	// Quantity of dataset items which not produce overflow of inspected function
	NotOverflowedItemsQuantity int
	// Quantity of dataset items which produce overflow of inspected function
	OverflowedItemsQuantity int
	// Function that returns a reference value
	Reference inspect.Reference
	// Writer associated with dataset storage
	Writer io.Writer
	// List of fillers that fill arguments of inspected and reference functions with
	// values. If not specified will be used filler.Boundary and filler.Rand fillers
	Fillers []filler.Filler[Type]

	// Minimum and maximum value for specified type
	min int64
	max int64

	// Arguments buffers, used to decrease allocations
	args   []Type
	args64 []int64

	// Dataset item buffer, used to decrease allocations
	item []byte
}

// Validates options. A reference function and writer must be specified.
func (clctr Collector[Type]) IsValid() error {
	if clctr.Reference == nil {
		return inspect.ErrReferenceNotSpecified
	}

	if clctr.Writer == nil {
		return ErrWriterNotSpecified
	}

	return nil
}

func (clctr Collector[Type]) normalize() Collector[Type] {
	if len(clctr.Fillers) == 0 {
		clctr.Fillers = append(clctr.Fillers, filler.NewBoundary[Type]())
		clctr.Fillers = append(clctr.Fillers, filler.NewRand[Type]())
	}

	return clctr
}

// Performs collecting dataset to file.
func (clctr Collector[Type]) CollectToFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}

	defer file.Close()

	clctr.Writer = file

	return clctr.Collect()
}

// Performs collecting.
func (clctr Collector[Type]) Collect() error {
	if err := clctr.IsValid(); err != nil {
		return err
	}

	clctr = clctr.normalize()

	clctr.min, clctr.max = inspect.PickUpRange[Type]()

	clctr.args = make([]Type, clctr.ArgsQuantity)
	clctr.args64 = make([]int64, clctr.ArgsQuantity)
	clctr.item = make([]byte, calcMaxItemLength(clctr.ArgsQuantity))

	return clctr.main()
}

func (clctr *Collector[Type]) main() error {
	for !clctr.isCollected() {
		if err := clctr.fillArgs(); err != nil {
			return err
		}

		if !clctr.isUseArgs() {
			continue
		}

		if err := clctr.writeItem(); err != nil {
			return err
		}
	}

	return nil
}

func (clctr *Collector[Type]) fillArgs() error {
	for _, filler := range clctr.Fillers {
		completed, err := filler.Fill(clctr.args, clctr.args64)
		if err != nil {
			return err
		}

		if completed {
			continue
		}

		return nil
	}

	return ErrNotEnoughDataInFillers
}

func (clctr *Collector[Type]) isCollected() bool {
	return clctr.OverflowedItemsQuantity <= 0 && clctr.NotOverflowedItemsQuantity <= 0
}

func (clctr *Collector[Type]) isUseArgs() bool {
	reference, _ := clctr.Reference(clctr.args64...)

	if reference > clctr.max || reference < clctr.min {
		if clctr.OverflowedItemsQuantity <= 0 {
			return false
		}

		clctr.OverflowedItemsQuantity--

		return true
	}

	if clctr.NotOverflowedItemsQuantity <= 0 {
		return false
	}

	clctr.NotOverflowedItemsQuantity--

	return true
}

func (clctr *Collector[Type]) writeItem() error {
	return writeItem(clctr.Writer, clctr.item, clctr.Reference, clctr.args, clctr.args64)
}

// Writes dataset item to specified writer.
func WriteItem[Type inspect.EightBits](
	writer io.Writer,
	reference inspect.Reference,
	args ...Type,
) error {
	buffer := make([]byte, calcMaxItemLength(len(args)))

	args64 := make([]int64, len(args))

	for id := range args {
		args64[id] = int64(args[id])
	}

	return writeItem(writer, buffer, reference, args, args64)
}

func writeItem[Type inspect.EightBits](
	writer io.Writer,
	buffer []byte,
	reference inspect.Reference,
	args []Type,
	args64 []int64,
) error {
	buffer = buffer[:0]

	ref, fault := reference(args64...)

	buffer = strconv.AppendBool(buffer, fault != nil)

	buffer = append(buffer, " "...)
	buffer = strconv.AppendInt(buffer, ref, consts.DecimalBase)

	for _, arg := range args {
		buffer = append(buffer, " "...)
		buffer = strconv.AppendInt(buffer, int64(arg), consts.DecimalBase)
	}

	buffer = append(buffer, '\n')

	if _, err := writer.Write(buffer); err != nil {
		return err
	}

	return nil
}

func calcMaxItemLength(argsQuantity int) int {
	const (
		maxFaultLen     = max(len("false"), len("true"))
		maxReferenceLen = len(" 18446744073709551615")
		maxArgLen       = len(" -127")
		maxNewLineLen   = len("\n")
	)

	return maxFaultLen + maxReferenceLen + maxArgLen*argsQuantity + maxNewLineLen
}
