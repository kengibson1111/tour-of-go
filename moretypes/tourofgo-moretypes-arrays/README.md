# moretypes - arrays

This is a simple example creating an array on the stack. arrays have a hard limit - you specify the length and the array is only that length because the length is included in the type at runtime. In other words, the length is embedded in the type, so [4]int and [5]int are actually different types. Arrays are values, not pointers.

Thoughts on the Go string implementation: just an assumption, but it looks like a string is a pointer to a byte array. Memory for the byte array is not initialized until the string is initialized with a value. Then when a string is initialized (like in the sample code), byte array memory is allocated and set.

Therefore, the sample code creates an array of 2 byte-array pointers on the stack. Then Go allocates memory for the byte arrays on the stack or heap depending on the underlying implementation. This seems like a small point, but when you assign and reassign string (byte-array pointer) values this could be expensive for the runtime. Go hides all of the string complexity in the language and runtime, but the complexity has to be addressed eventually.
