package dataset

import (
	"io"
	"maps"
	"os"
	"strconv"

	"github.com/akramarenkov/safe/internal/consts"
	"github.com/akramarenkov/safe/internal/inspect"
	"github.com/akramarenkov/safe/internal/inspect/dataset/filler"
	"github.com/akramarenkov/safe/internal/inspect/types"
	"github.com/akramarenkov/safe/intspec"
)

// Options of collecting. A reference function and writer must be specified.
type Collector[Type types.UpToUSI32] struct {
	// Quantity of arguments for inspected and reference functions
	ArgsQuantity int
	// Quantity of dataset items which not produce overflow of inspected function
	NotOverflowedItemsQuantity int
	// Quantity of dataset items which produce overflow of inspected function
	OverflowedItemsQuantity int
	// Function that returns a reference value
	types.Reference[int64]
	// Quantity limits for reference values
	ReferenceLimits map[int64]uint
	// Writer associated with dataset storage
	Writer io.Writer
	// List of fillers that fill arguments of dataset item (reference function) by
	// values. If not specified will be used filler [filler.Set] with
	// [filler.Boundaries] setter and filler [filler.Rand]
	Fillers []filler.Filler[Type]

	// Minimum and maximum value for specified type
	minimum int64
	maximum int64

	// Arguments buffers, used to decrease allocations
	args      []Type
	args64    []int64
	args64Dup []int64

	// Dataset item buffer, used to decrease allocations
	item []byte

	// map used to maintain uniqueness of items in a dataset
	unique map[string]struct{}
}

// Validates options. A reference function and writer must be specified.
func (cltr Collector[Type]) IsValid() error {
	if cltr.Reference == nil {
		return inspect.ErrReferenceNotSpecified
	}

	if cltr.Writer == nil {
		return ErrWriterNotSpecified
	}

	return nil
}

func (cltr Collector[Type]) normalize() Collector[Type] {
	if len(cltr.Fillers) == 0 {
		cltr.Fillers = append(cltr.Fillers, filler.NewSet[Type](), filler.NewRand[Type]())
	}

	return cltr
}

// Performs collecting dataset to file.
func (cltr Collector[Type]) CollectToFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}

	defer file.Close()

	cltr.Writer = file

	return cltr.Collect()
}

// Performs collecting dataset.
func (cltr Collector[Type]) Collect() error {
	if err := cltr.IsValid(); err != nil {
		return err
	}

	cltr = cltr.normalize()

	cltr.ReferenceLimits = maps.Clone(cltr.ReferenceLimits)

	cltr.minimum, cltr.maximum = inspect.ConvSpan[Type, int64]()

	cltr.args = make([]Type, cltr.ArgsQuantity)
	cltr.args64 = make([]int64, cltr.ArgsQuantity)
	cltr.args64Dup = make([]int64, cltr.ArgsQuantity)
	cltr.item = make([]byte, calcMaxItemLength[Type](cltr.ArgsQuantity))

	cltr.unique = make(map[string]struct{}, cltr.calcDatasetLength())

	return cltr.main()
}

func (cltr Collector[Type]) calcDatasetLength() int {
	length := 0

	if cltr.NotOverflowedItemsQuantity > 0 {
		length += cltr.NotOverflowedItemsQuantity
	}

	if cltr.OverflowedItemsQuantity > 0 {
		length += cltr.OverflowedItemsQuantity
	}

	return length
}

func (cltr *Collector[Type]) main() error {
	for !cltr.isCollected() {
		if err := cltr.fillArgs(); err != nil {
			return err
		}

		if !cltr.isUseArgs() {
			continue
		}

		if err := cltr.writeItem(); err != nil {
			return err
		}
	}

	return nil
}

func (cltr *Collector[Type]) isCollected() bool {
	return cltr.OverflowedItemsQuantity <= 0 && cltr.NotOverflowedItemsQuantity <= 0
}

func (cltr *Collector[Type]) fillArgs() error {
	for _, flr := range cltr.Fillers {
		completed, err := flr.Fill(cltr.args, cltr.args64)
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

func (cltr *Collector[Type]) isUseArgs() bool {
	reference, fault := cltr.Reference(cltr.dupArgs64()...)

	if cltr.isLimited(reference) {
		return false
	}

	if reference > cltr.maximum || reference < cltr.minimum {
		if cltr.OverflowedItemsQuantity <= 0 {
			return false
		}

		cltr.prepareItem(reference, fault)

		if !cltr.isUnique() {
			return false
		}

		cltr.OverflowedItemsQuantity--

		return true
	}

	if cltr.NotOverflowedItemsQuantity <= 0 {
		return false
	}

	cltr.prepareItem(reference, fault)

	if !cltr.isUnique() {
		return false
	}

	cltr.NotOverflowedItemsQuantity--

	return true
}

func (cltr *Collector[Type]) isLimited(reference int64) bool {
	if len(cltr.ReferenceLimits) == 0 {
		return false
	}

	quantity, exists := cltr.ReferenceLimits[reference]
	if !exists {
		return false
	}

	if quantity == 0 {
		return true
	}

	cltr.ReferenceLimits[reference]--

	return false
}

func (cltr *Collector[Type]) isUnique() bool {
	key := string(cltr.item)

	if _, exists := cltr.unique[key]; exists {
		return false
	}

	cltr.unique[key] = struct{}{}

	return true
}

func (cltr *Collector[Type]) prepareItem(reference int64, fault error) {
	cltr.item = prepareItem(cltr.item, reference, fault, cltr.args...)
}

func (cltr *Collector[Type]) writeItem() error {
	if _, err := cltr.Writer.Write(cltr.item); err != nil {
		return err
	}

	return nil
}

// Protection against changes args64 from the reference function.
func (cltr *Collector[Type]) dupArgs64() []int64 {
	copy(cltr.args64Dup, cltr.args64)
	return cltr.args64Dup
}

// Writes dataset item to specified writer.
func WriteItem[Type types.UpToUSI32](
	writer io.Writer,
	reference types.Reference[int64],
	args ...Type,
) error {
	buffer := make([]byte, calcMaxItemLength[Type](len(args)))

	args64 := make([]int64, len(args))

	for id := range args {
		args64[id] = int64(args[id])
	}

	ref, fault := reference(args64...)

	buffer = prepareItem(buffer, ref, fault, args...)

	if _, err := writer.Write(buffer); err != nil {
		return err
	}

	return nil
}

func prepareItem[Type types.UpToUSI32](
	buffer []byte,
	reference int64,
	fault error,
	args ...Type,
) []byte {
	buffer = buffer[:0]

	buffer = strconv.AppendBool(buffer, fault != nil)

	buffer = append(buffer, " "...)
	buffer = strconv.AppendInt(buffer, reference, consts.DecimalBase)

	for _, arg := range args {
		buffer = append(buffer, " "...)
		buffer = strconv.AppendInt(buffer, int64(arg), consts.DecimalBase)
	}

	buffer = append(buffer, '\n')

	return buffer
}

func calcMaxItemLength[Type types.UpToUSI32](argsQuantity int) int {
	const (
		maxFaultLen        = len("false")
		maxReferenceLen    = len(" -9223372036854775808")
		maxArgSeparatorLen = len(" ")
		maxNewLineLen      = len("\n")
	)

	minimum, maximum := intspec.Range[Type]()

	maxArgLen := max(
		len(strconv.FormatInt(int64(minimum), consts.DecimalBase)),
		len(strconv.FormatInt(int64(maximum), consts.DecimalBase)),
	)

	length := maxFaultLen +
		maxReferenceLen +
		maxArgSeparatorLen*argsQuantity +
		maxArgLen*argsQuantity +
		maxNewLineLen

	return length
}
