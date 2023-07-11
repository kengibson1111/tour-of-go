# moretypes - slices (default)

This shows how defaults are applied for the slice low-bound index and length when no value is specified. Slices can be created from arrays or from other slices. When a slice is created from an array, the default low-bound index for the slice is zero. The default length is the array length.

When a slice is created from another slice, the default low-bound index is zero (which correlates to the other slice and not the array). The default length is the length of the other slice and not the length of the array. The new slice's array pointer is the same as the other slice's array pointer.
