# moretypes - slices

This shows how effective pointers are since a slice has a pointer to an array, a low-bound index into the array, and a high-bound index into the array. That's it. Slices do not store data like an array.

The sample code creates an array of ints using literals. Then the slice to the int array is created. The slice has the pointer. I would think that if a function needs an array as a parameter, you want to use a slice for that because that is the best optimization of the function call stack.
