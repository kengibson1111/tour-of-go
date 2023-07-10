# moretypes - slices

This shows how effective pointers are since a slice has a pointer to an array, a low-bound index into the array, and a high-bound index into the array. That's it. Slices do not store data like an array. The slice has access to elements in the array from the low-bound index to the (high-bound index - 1). In other words, the high-bound index itself is excluded from the slice.

The sample code creates an array of ints using literals. Then the slice to the int array is created. The slice has the pointer. I would think that if a function needs an array as a parameter, you want to use a slice for that because that is the best optimization of the function call stack.
