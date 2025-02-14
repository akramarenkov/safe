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

	"github.com/akramarenkov/intspec"
)

// Options of collecting. A reference function must be specified.
type Opts[Type types.UpToUSI32] struct {
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
	// List of fillers that fill arguments of dataset item (reference function) by
	// values. If not specified will be used filler [filler.Set] with
	// [filler.Boundaries] setter and filler [filler.Rand]
	Fillers []filler.Filler[Type]
}

func (opts Opts[Type]) isValid() error {
	if opts.Reference == nil {
		return inspect.ErrReferenceNotSpecified
	}

	return nil
}

func (opts Opts[Type]) normalize() Opts[Type] {
	if len(opts.Fillers) == 0 {
		opts.Fillers = []filler.Filler[Type]{filler.NewSet[Type](), filler.NewRand[Type]()}
	}

	return opts
}

func (opts Opts[Type]) datasetLength() int {
	length := 0

	if opts.NotOverflowedItemsQuantity > 0 {
		length += opts.NotOverflowedItemsQuantity
	}

	if opts.OverflowedItemsQuantity > 0 {
		length += opts.OverflowedItemsQuantity
	}

	return length
}

type collector[Type types.UpToUSI32] struct {
	opts Opts[Type]

	// Writer associated with dataset storage
	writer io.Writer

	// Quantity limits for reference values
	referenceLimits map[int64]uint

	// Quantity of dataset items which not produce overflow of inspected function
	notOverflowedItemsQuantity int
	// Quantity of dataset items which produce overflow of inspected function
	overflowedItemsQuantity int

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

// Performs collecting dataset to file.
func CollectToFile[Type types.UpToUSI32](opts Opts[Type], path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, defaultFileMode)
	if err != nil {
		return err
	}

	defer file.Close()

	return Collect(opts, file)
}

// Performs collecting dataset to writer.
func Collect[Type types.UpToUSI32](opts Opts[Type], writer io.Writer) error {
	if err := opts.isValid(); err != nil {
		return err
	}

	if writer == nil {
		return ErrWriterNotSpecified
	}

	opts = opts.normalize()

	minimum, maximum := intspec.Range[Type]()

	cltr := &collector[Type]{
		opts: opts,

		writer: writer,

		referenceLimits:            maps.Clone(opts.ReferenceLimits),
		notOverflowedItemsQuantity: opts.NotOverflowedItemsQuantity,
		overflowedItemsQuantity:    opts.OverflowedItemsQuantity,

		minimum: int64(minimum),
		maximum: int64(maximum),

		args:      make([]Type, opts.ArgsQuantity),
		args64:    make([]int64, opts.ArgsQuantity),
		args64Dup: make([]int64, opts.ArgsQuantity),
		item:      make([]byte, calcMaxItemLength[Type](opts.ArgsQuantity)),
		unique:    make(map[string]struct{}, opts.datasetLength()),
	}

	return cltr.main()
}

func (cltr *collector[Type]) main() error {
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

func (cltr *collector[Type]) isCollected() bool {
	return cltr.overflowedItemsQuantity <= 0 && cltr.notOverflowedItemsQuantity <= 0
}

func (cltr *collector[Type]) fillArgs() error {
	for _, flr := range cltr.opts.Fillers {
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

func (cltr *collector[Type]) isUseArgs() bool {
	reference, fault := cltr.opts.Reference(cltr.dupArgs64()...)

	if cltr.isLimited(reference) {
		return false
	}

	if reference > cltr.maximum || reference < cltr.minimum {
		if cltr.overflowedItemsQuantity <= 0 {
			return false
		}

		cltr.prepareItem(reference, fault)

		if !cltr.isUnique() {
			return false
		}

		cltr.overflowedItemsQuantity--

		return true
	}

	if cltr.notOverflowedItemsQuantity <= 0 {
		return false
	}

	cltr.prepareItem(reference, fault)

	if !cltr.isUnique() {
		return false
	}

	cltr.notOverflowedItemsQuantity--

	return true
}

func (cltr *collector[Type]) isLimited(reference int64) bool {
	if len(cltr.referenceLimits) == 0 {
		return false
	}

	quantity, exists := cltr.referenceLimits[reference]
	if !exists {
		return false
	}

	if quantity == 0 {
		return true
	}

	cltr.referenceLimits[reference]--

	return false
}

func (cltr *collector[Type]) isUnique() bool {
	key := string(cltr.item)

	if _, exists := cltr.unique[key]; exists {
		return false
	}

	cltr.unique[key] = struct{}{}

	return true
}

func (cltr *collector[Type]) prepareItem(reference int64, fault error) {
	cltr.item = prepareItem(cltr.item, reference, fault, cltr.args...)
}

func (cltr *collector[Type]) writeItem() error {
	if _, err := cltr.writer.Write(cltr.item); err != nil {
		return err
	}

	return nil
}

// Protection against changes args64 from the reference function.
func (cltr *collector[Type]) dupArgs64() []int64 {
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
