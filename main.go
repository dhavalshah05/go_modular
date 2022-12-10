package main

import (
	"fmt"
	"sort"
)

func main() {

	names := []string{"A", "B", "C"}
	fmt.Println("Real:", names)

	// Add at end
	atEnd := append(names, "P")
	fmt.Println("Add at end:", atEnd)

	// Add at start
	atStart := append([]string{"P"}, names...)
	fmt.Println("Add at start:", atStart)

	// Add at index
	addIndex := 2
	var addAtIndex []string
	if len(names) <= addIndex {
		addAtIndex = append(names, "P")
	} else {
		addAtIndex = append(addAtIndex, names[0:addIndex]...)
		addAtIndex = append(addAtIndex, "P")
		addAtIndex = append(addAtIndex, names[addIndex:]...)
	}
	fmt.Printf("Add at index %d: %v\n", addIndex, addAtIndex)

	// Remove from end
	removeFromEnd := append([]string{}, names[:len(names)-1]...)
	fmt.Println("Remove from end:", removeFromEnd)

	// Remove from start
	removeFromStart := append([]string{}, names[1:]...)
	fmt.Println("Remove from start:", removeFromStart)

	// Remove from index
	removeIndex := 1
	if removeIndex >= len(names) {
		fmt.Printf("Cannot remove value from index %d\n", removeIndex)
	} else {
		removeFromIndex := append([]string{}, names[:removeIndex]...)
		removeFromIndex = append(removeFromIndex, names[removeIndex+1:]...)
		fmt.Printf("Remove from index %d: %v\n", removeIndex, removeFromIndex)
	}

	// Combine 2 slices
	slice1 := []string{"1", "2"}
	slice2 := []string{"3", "4"}
	combineSlices := append(slice1, slice2...)
	fmt.Println("Combine slices: ", combineSlices)

	// Sort slice
	unsortedSlice := []int{1, 4, 2, 7}
	sort.Slice(unsortedSlice, func(i, j int) bool {
		return unsortedSlice[i] < unsortedSlice[j]
	})
	fmt.Printf("Sorted: %v\n", unsortedSlice)

	// Filter slice
	unfilteredSlice := []int{1, 4, 2, 7}
	var filteredSlice []int
	for _, number := range unfilteredSlice {
		if number%2 == 0 {
			filteredSlice = append(filteredSlice, number)
		}
	}
	fmt.Printf("Filtered: %v\n", filteredSlice)

	// Map slice
	normalSlice := []int{1, 4, 2, 7}
	var mappedSlice []int
	for _, number := range normalSlice {
		mappedSlice = append(mappedSlice, number*10)
	}
	fmt.Printf("Mapped: %v\n", mappedSlice)
}
