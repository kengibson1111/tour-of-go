# methods - interfaces (assert)

Types can be asserted and the assertion returns the value and a bool that indicates if the assertion worked. Just the value can be returned, but that will cause a runtime error if the interface type does not match the assertion. Therefore, it is a best practice to return the value and bool from an assertion if there is any chance that the interface type does not match the assertion.
