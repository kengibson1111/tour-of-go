# moretypes - struct pointers

This is showing that a pointer to a struct can be used to make changes in the struct instance. C makes a field access distinction between a struct and a pointer. Go doesn't do that - dot notation all the way. In this sample, it looks like a a struct instance is created on the stack and the pointer value is stack address space.
