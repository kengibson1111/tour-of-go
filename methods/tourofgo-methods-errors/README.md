# methods - errors

Another interface from the core package fmt is `error`. Many Go functions and methods return an instance of an error implementation. Then the caller can check if the returned error is nil and react appropriately. The sample code shows how to implement the error interface.

In `run()`, MyError is created and the pointer is returned which aligns with the pointer receiver of the Error() method defined in the error interface. So the runtime knows MyError implements the error interface and binds it. Now MyError's Error() method can be called through the error interface by any other function. In this case, Println can call it.
