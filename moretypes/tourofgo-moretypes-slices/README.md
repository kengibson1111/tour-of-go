# moretypes - slices

This shows how effective pointers are since a slice has a pointer to an array, a low-bound index into the array, a length, and a capacity. Slices do not store data like an array.

Slice expressions have a low-bound index and a high-bound index. Slices can be created from arrays or from other slices. When a slice is created from an array, the slice expression's default low-bound index is zero. The slice expression's default high-bound index is the array length. The array element related to the slice expression's high-bound index is not included in the new slice.

The sample code creates an array of ints using literals. Then the slice to the int array is created. I would think that if a function needs an array as a parameter, it's more optimized to use a slice for that because of the small slice footprint for the function call stack.
