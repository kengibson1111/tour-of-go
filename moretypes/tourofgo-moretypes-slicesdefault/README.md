# moretypes - slices (default)

This shows how defaults are applied for slice expressions when a low-bound index or high-bound index is not specified. Slices can be created from arrays or from other slices. When a slice is created from an array, the slice expression's default low-bound index is zero. The slice expression's default high-bound index is the array length. Remember, the array element related to the slice expression's high-bound index is not included in the new slice.

When a slice is created from another slice, the new slice expression's default low-bound index is zero (which correlates to the other slice and not the array). The new slice expression's default high-bound index is the length of the other slice and not the length of the array. The new slice's array pointer is the same as the other slice's array pointer.

The code shows how the slice reference struct keeps the same array literal pointer when "resliced", but the length and capacity changes after each "reslice". I expanded the fmt.Println statements to show this and added comments. Slice length and capacity are described more in **moretypes - slices (length and capacity)**.
