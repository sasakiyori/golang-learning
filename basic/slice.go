package main
import "fmt"

// test for slice: a structure or an action
func main() {  
	// slice structure
	simple_slice := make([]int, 3, 10)
	fmt.Println(simple_slice)
	fmt.Println(len(simple_slice))	// length = 3
	fmt.Println(cap(simple_slice))	// capacity = 10
	println(simple_slice)      // [length/capacity]pointer-address
	fmt.Println(simple_slice[0])

	// difference between slice and array
	instantiated_slice := []string{"a", "b", "c"}
	instantiated_array := [3]string{"a", "b", "c"}
	fmt.Printf("instantiated_slice: %s, instantiated_array: %s\n", instantiated_slice, instantiated_array)

	// slice action
	// sliced_child only includes index 0, 1 from instantiated_slice
	sliced_child := instantiated_slice[0:2]
	fmt.Printf("after slice, instantiated_slice: %s, sliced_child: %s\n", instantiated_slice, sliced_child)


	// change the base slice's element, the element of sliced_child will change as well
	instantiated_slice[1] = "change"
	fmt.Printf("after change, instantiated_slice: %s, sliced_child: %s\n", instantiated_slice, sliced_child)

	// append by child slice which's next element is in range of base slice, all elements at that index will change
	child_appended_slice := append(sliced_child, "child_append")
	fmt.Printf("after chiled_append, instantiated_slice: %s, sliced_child: %s, child_appended_slice: %s\n",
				instantiated_slice, sliced_child, child_appended_slice)

	// append by base slice, both base slice and clild slice's element will not change
	base_appended_slice := append(instantiated_slice, "base_append")
	fmt.Printf("after base_append, instantiated_slice: %s, sliced_child: %s, child_appended_slice: %s, base_appended_slice: %s\n",
				instantiated_slice, sliced_child, child_appended_slice, base_appended_slice)

	// append by child slice which's next element is out of range of base slice, all elements at that index will not change
	child_appended_slice2 := append(child_appended_slice, "child_append_2")
	fmt.Printf("after base_append, instantiated_slice: %s, sliced_child: %s, child_appended_slice: %s, base_appended_slice: %s, child_appended_slice2: %s\n",
				instantiated_slice, sliced_child, child_appended_slice, base_appended_slice, child_appended_slice2)

	// append by child slice which's next element is out of range of base slice, all elements at that index will not change
	sliced_child_2 := instantiated_slice[1:]
	sliced_child_2_append := append(sliced_child_2, "append_end")
	fmt.Printf("after append end, instantiated_slice: %s, sliced_child_2: %s, sliced_child_2_append: %s\n",
				instantiated_slice, sliced_child_2, sliced_child_2_append)


	
	// after copying, the new slice doesn't share the same pointer/address with the base slice
	slice_a := []string{"1", "2", "3"}
	slice_b := []string{"5", "6", "7", "8", "9"}
	num := copy(slice_a, slice_b)
	fmt.Println(num)
	fmt.Printf("after copy, slice_a: %s, slice_b: %s\n", slice_a, slice_b)
	slice_b[1] = "change"
	fmt.Printf("after change, slice_a: %s, slice_b: %s\n", slice_a, slice_b)
}