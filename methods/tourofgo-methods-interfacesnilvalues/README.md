# methods - interfaces (nil values)

In **methods - interfaces (value)**, suppose t was nil when i was assigned. In Java, that would mean a null interface and a potential null pointer exception. In Go, the interface is a tuple of a value and the type of that value. It has a nil concrete value and type before assignment and is considered a nil interface.

After assignment, the type is not nil *even if the related value is nil*. The interface is bound to a matching method with the right receiver for that type. The concrete value could be nil, but in Go the method implementation can gracefully check for a nil concrete value.
