package writer

//
//import (
//	"reflect"
//	"sync"
//)
//
//// pp is used to store a printer's state and is reused with sync.Pool to avoid allocations.
//type pp struct {
//	buf buffer
//
//	// arg holds the current item, as an interface{}.
//	arg any
//
//	// value is used instead of arg for reflect values.
//	value reflect.Value
//
//	// fmt is used to format basic items such as integers or strings.
//	fmt fmt
//
//	// reordered records whether the format string used argument reordering.
//	reordered bool
//	// goodArgNum records whether the most recent reordering directive was valid.
//	goodArgNum bool
//	// panicking is set by catchPanic to avoid infinite panic, recover, panic, ... recursion.
//	panicking bool
//	// erroring is set when printing an error string to guard against calling handleMethods.
//	erroring bool
//	// wrapErrs is set when the format string may contain a %w verb.
//	wrapErrs bool
//	// wrappedErrs records the targets of the %w verb.
//	wrappedErrs []int
//}
//
//var ppFree = sync.Pool{
//	New: func() any { return new(pp) },
