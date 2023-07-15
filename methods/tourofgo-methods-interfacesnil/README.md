# methods - interfaces (nil)

We can still get into trouble calling a method signature on an interface that has not been assigned yet. It causes a runtime error because a method with the correct receiver has not been bound yet. Use a `== nil` check before calling any interface method.
