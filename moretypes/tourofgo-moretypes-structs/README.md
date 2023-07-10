# moretypes - structs

In Go, a struct is a collection of fields in its most basic form. For this lesson, I assume that the struct instance is being created on the stack (not heap). I also assume fmt.Println uses a copy of the struct instance (function stack to call stack) instead of a pointer to the struct instance for the call stack.
