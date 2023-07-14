# methods - methods (receiver)

This verifies a Go "best practice" for receivers. For a given type, receivers for all methods should be either value or pointer but not a mix of both. And as mentioned in **methods - methods (pointers)**, you do get better function call stack performance using a pointer receiver for methods that only read if the type is a large struct.
