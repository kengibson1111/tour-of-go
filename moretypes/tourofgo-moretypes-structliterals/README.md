# moretypes - struct literals

This shows a few examples of creating a struct instance using literals. In the pointer case, the struct instance is created first and then the pointer is assigned.

And, I assume for the fmt.Println call at runtime, 3 struct instances are being copied from the stack to the function call stack and 1 pointer value is being copied from the stack to the function call stack. Doesn't seem like a big deal when a struct is so small, but it could be a big deal with more complex structs.
