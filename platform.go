package main

import "fmt"

// +gen  slice:"Where,Count,GroupBy[string],Aggregate[string]"
type Platform struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	Config  map[string]interface{} `json:"config"`
}

func (p Platform) String() string {
	return fmt.Sprintf("this is my platform print: %s , %s", p.Name, p.Version)
}
