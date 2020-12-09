// Generated by: main
// TypeWriter: slice
// Directive: +gen on Platform

package main

// PlatformSlice is a slice of type Platform. Use it where you would use []Platform.
type PlatformSlice []Platform

// Where returns a new PlatformSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv PlatformSlice) Where(fn func(Platform) bool) (result PlatformSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of PlatformSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv PlatformSlice) Count(fn func(Platform) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv PlatformSlice) GroupByString(fn func(Platform) string) map[string]PlatformSlice {
	result := make(map[string]PlatformSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// AggregateString iterates over PlatformSlice, operating on each element while maintaining ‘state’. See: http://clipperhouse.github.io/gen/#Aggregate
func (rcv PlatformSlice) AggregateString(fn func(string, Platform) string) (result string) {
	for _, v := range rcv {
		result = fn(result, v)
	}
	return
}
