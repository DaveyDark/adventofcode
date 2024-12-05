package day05

import (
	"reflect"

	"github.com/emirpasic/gods/sets/hashset"
)

func verifyUpdate(update []int, rules map[int]hashset.Set) bool {
	// Build set of update
	updateSet := hashset.New()
	for _, up := range update {
		updateSet.Add(up)
	}
	// Verify all updates
	prev := hashset.New()
	for _, up := range update {
		// Check if the current element has rules defined
		rule, ok := rules[up]
		if ok {
			// Check if it follows defined rules
			iRule := rule.Intersection(updateSet)
			if !reflect.DeepEqual(prev.Intersection(iRule), iRule) {
				return false
			}
		}
		prev.Add(up)
	}
	return true
}
