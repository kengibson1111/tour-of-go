# methods - methods (pointers)

If you want to modify a type instance, the receiver for the method has to be a pointer to the type. So anything that reads from a type instance could have just a value receiver. The only argument against that is function call stack efficiency since having a value receiver means the entire type instance *value* will be copied to the call stack for the method. In the sample code, change the pointer receiver to a value receiver and see what happens. Notice that dot notation is used to access methods with both types of receivers.
