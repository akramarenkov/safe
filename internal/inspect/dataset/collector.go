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
	"github.com/akramarenkov/safe/internal/intspan"
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
	Reference
	// Quantity limits for reference values
	ReferenceLimits map[int64]uint
	// Writer associated with dataset storage
	Writer io.Writer
	// List of fillers that fill arguments of dataset item (reference function) by
	// values. If not specified will be used filler [filler.Set] with
	// [filler.Boundaries] setter and filler [filler.Rand]
	Fillers []filler.Filler[Type]

	// Minimum and maximum value for specified type
	min int64
	max int64

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
		clctr.Fillers = append(clctr.Fillers, filler.NewSet[Type]())
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

	clctr.ReferenceLimits = maps.Clone(clctr.ReferenceLimits)

	clctr.min, clctr.max = inspect.ConvSpan[Type, int64]()

	clctr.args = make([]Type, clctr.ArgsQuantity)
	clctr.args64 = make([]int64, clctr.ArgsQuantity)
	clctr.args64Dup = make([]int64, clctr.ArgsQuantity)
	clctr.item = make([]byte, calcMaxItemLength[Type](clctr.ArgsQuantity))

	clctr.unique = make(map[string]struct{}, clctr.calcDatasetLength())

	return clctr.main()
}

func (clctr Collector[Type]) calcDatasetLength() int {
	length := 0

	if clctr.NotOverflowedItemsQuantity > 0 {
		length += clctr.NotOverflowedItemsQuantity
	}

	if clctr.OverflowedItemsQuantity > 0 {
		length += clctr.OverflowedItemsQuantity
	}

	return length
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

func (clctr *Collector[Type]) isCollected() bool {
	return clctr.OverflowedItemsQuantity <= 0 && clctr.NotOverflowedItemsQuantity <= 0
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

func (clctr *Collector[Type]) isUseArgs() bool {
	reference, fault := clctr.Reference(clctr.dupArgs64()...)

	if clctr.isLimited(reference) {
		return false
	}

	if reference > clctr.max || reference < clctr.min {
		if clctr.OverflowedItemsQuantity <= 0 {
			return false
		}

		clctr.prepareItem(reference, fault)

		if !clctr.isUnique() {
			return false
		}

		clctr.OverflowedItemsQuantity--

		return true
	}

	if clctr.NotOverflowedItemsQuantity <= 0 {
		return false
	}

	clctr.prepareItem(reference, fault)

	if !clctr.isUnique() {
		return false
	}

	clctr.NotOverflowedItemsQuantity--

	return true
}

func (clctr *Collector[Type]) isLimited(reference int64) bool {
	if len(clctr.ReferenceLimits) == 0 {
		return false
	}

	quantity, exists := clctr.ReferenceLimits[reference]
	if !exists {
		return false
	}

	if quantity == 0 {
		return true
	}

	clctr.ReferenceLimits[reference]--

	return false
}

func (clctr *Collector[Type]) isUnique() bool {
	key := string(clctr.item)

	if _, exists := clctr.unique[key]; exists {
		return false
	}

	clctr.unique[key] = struct{}{}

	return true
}

func (clctr *Collector[Type]) prepareItem(reference int64, fault error) {
	clctr.item = prepareItem(clctr.item, reference, fault, clctr.args...)
}

func (clctr *Collector[Type]) writeItem() error {
	if _, err := clctr.Writer.Write(clctr.item); err != nil {
		return err
	}

	return nil
}

// Protection against changes args64 from the reference function.
func (clctr *Collector[Type]) dupArgs64() []int64 {
	copy(clctr.args64Dup, clctr.args64)
	return clctr.args64Dup
}

// Writes dataset item to specified writer.
func WriteItem[Type types.USI8](
	writer io.Writer,
	reference Reference,
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

	min, _, _ := intspan.Get[Type]()

	maxArgLen := len(strconv.FormatInt(int64(min), consts.DecimalBase))

	length := maxFaultLen +
		maxReferenceLen +
		maxArgSeparatorLen*argsQuantity +
		maxArgLen*argsQuantity +
		maxNewLineLen

	return length
}
