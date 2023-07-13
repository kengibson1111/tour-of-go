# methods - methods (non-struct)

This is showing a few things. Types don't necessarily have to be defined for structs. The sample code defines a type for a primitive. Why? For this sample, it is a way to overload functionality in a float64. Define a type and then define a method for the type. Methods can only be defined for types in the same package, and the type cannot be a pointer to another type. So in the sample, there is no way to define a method for a float64, but there is a way to define a type to a float64 and then define a method for the custom type.
