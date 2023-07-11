# moretypes - slices (length and capacity)

This prints out length and capacity of a slice. A slice is a reference struct to an array, so it contains a pointer to an array, a low-bound array index, a length which describes how many more elements from array[low-bound index] to include in the slice, and a capacity.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice (array[low-bound index]). In other words, the slice's capacity is the array length minus the slice's low-bound index. So, if the underlying array has 10 elements in it and the slice's low-bound index is 2, the slice's capacity is 8.

The code also shows how the slice reference struct keeps the same array literal pointer when "resliced".