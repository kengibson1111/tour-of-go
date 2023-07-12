# tour-of-go ![tour-of-go](/images-repo/icons8-golang-240.png)

This is a companion repo for Go's [Tour of Go](https://go.dev/tour/welcome/1).[^1] Definitely use the code window on the tour. This repo contains mostly code from the tour, and at times I add more code based on ideas from the lessons. Code from the welcome pages is not included.

## Lessons

* range (skip) - shows how to skip the index if you just want a range value and how to skip the value if
  you just want the index.

* EXERCISE (slices) - this was fun to do. After finishing and comparing with other answers across github,
  it isn't the best code, but it works.

* maps - basic example of map use.

* maps (literals) - literals can be used for maps - like structs.

* maps (literals inference) - same as last example, but demonstrating type inference in the literal.

* maps (mutate) - basic CRUD for a map.

* EXERCISE (maps) - maybe not the best, but it was quick :).

* functions (value) - yes, functions are 1st class citizens in golang. And when that happens, the possibilities
  are almost endless. In the sample, compute() has a function parameter. All compute() knows is that the
  function will accept 2 values and compute() will pass 3 and 4 as the values to the function parameter. Now
  you can pass any function that meets the parameter spec and it will do something with the values 3 and 4.
  math.Pow meets the spec, but you can also define your own function that also meets the spec. compute()
  doesn't care.

* functions (closure) - closures are cool. They can reference variables outside of their local function scope.
  So in the sample, adder() is returning a closure that can mutate its own sum variable. Seems like a potential
  implementation for basic map reductions.

* EXERCISE (functions) - not an optimized implementation, but I wanted to do something with slices.

* methods - golang does not have objects, but it does have types. Types can be associated with primitives, structs,
  etc. If you add a type receiver argument in a function definition between the keyword func and the name of the
  function, you have now associated a type to a function. And that is golang's version of a method.

* methods (func) - this is just emphasizing that methods are functions. The sample code is the same as the previous
  lesson without the type receiver argument.

* methods (nonstruct) - this is showing a few things. Types don't necessarily have to be defined for structs. The
  sample code defines a type for a primitive. Why? For this sample, it is a way to overload functionality in a float64.
  Define a type and then define a method for the type. Methods can only be defined for types in the same package, and
  the type cannot be a pointer to another type. So in the sample, there is no way to define a method for a float64, but
  there is a way to define a type to a float64 and then define a method for the custom type.

* methods (pointers) - if you want to modify a type instance, the receiver for the method has to be a pointer to
  the type. So anything that reads from a type instance could have just a value receiver. The only argument
  against that is function call stack efficiency since having a value receiver means the entire type instance VALUE
  will be copied to the call stack for the method. In the sample code, change the pointer receiver to a value receiver and
  see what happens. SIDE NOTE: notice that dot notation is used to access methods with both types of receivers.

* functions (pointers) - just to drive home the concept, methods are functions. So the same arguments apply to passing
  a value of a type instance to prevent updates and a pointer of the type instance to allow updates. In the last
  lesson, you changed the pointer receiver to a value receiver. In this sample, change the *Vertex to Vertex in Scale()
  and you should see a compile error. Fix it and try again. This is one area where methods may have an advantage over
  simple functions.

* methods (pointersind) - emphasizing the indirection flexibility with methods. Functions that have a pointer
  parameter must take a pointer. Methods that have a pointer receiver can take a value or a pointer. golang figures
  it out.

* methods (pointersindrev) - and the indirection flexibility with methods is true in reverse, too. Functions
  that have a value parameter must take a value. Methods that have a value receiver can take a pointer or value.
  golang figures it out.

* methods (receiver) - this verifies a golang "good practice" for receivers. For a given type, receivers for all
  methods should be either value or pointer but not a mix of both. And as mentioned in the lesson "methods (pointers)"
  you do get better function call stack performance using a pointer receiver for methods that only read if the
  type is a large struct.

* interfaces - introduction of an interface type which is the same as in Java - a set of method signatures.

* interfaces (impl) - golang does not require explicit interface declaration like in Java. Just implement the
  method signature and it implicitly implements the interface.

* interfaces (value) - when an interface is declared, it has a nil value and type. It is considered a nil
  interface. The compiler checks to make sure each interface binding is legal. At runtime, the interface
  value and type are filled when the interface is assigned. Then one of the method signatures can be called
  exercising behavior based on the interface's value and type.

* interfaces (nilvalues) - this is where it starts getting interesting. In the previous lesson, suppose t was nil
  when i was assigned. In Java, that would mean a null interface and a potential null pointer exception. In golang,
  the interface is not nil after assignment. It has a nil concrete value and type before assignment and is a nil
  interface. After assignment, the type is not nil and the interface is bound to a matching method with the right
  receiver. The concrete value could be nil, but in golang the method implementation can gracefully check for a nil
  concrete value. Slick.

* interfaces (nil) - you can still get into trouble calling a method signature on an interface that has not been
  assigned yet. It causes a runtime error because a method with the correct receiver has not been bound yet.

* interfaces (empty) - just like in Java, you can have an empty interface. This does not bind any method signatures
  at runtime, but it does allow any type to be associated to an interface.

* interfaces (assert) - types can be asserted and the assertion returns the value and a bool that indicates if the
  assertion worked. Just the value can be returned, but that will cause a runtime error if the interface type
  does not match the assertion. So it is a good practice to return the value and bool from an assertion if
  there is any chance that the interface type does not match the assertion.

* interfaces (switch) - you can use a short assignment to a type assertion in a switch statement.

* stringer - Stringer is an interface from the core package fmt. The String() method signature is like the Java
  toString().

* EXERCISE (stringers) - pretty straightforward.

* errors - another interface from the core package fmt is error. Many golang functions and methods return an
  instance of an error implementation. Then the caller can check if the returned error is nil and react
  appropriately. The sample code shows how to implement the error interface. In run(), MyError is created
  and the pointer is returned which aligns with the pointer receiver of the Error() method defined in the
  error interface. So the runtime knows MyError implements the error interface and binds it. Now MyError's
  Error() method can be called through the error interface by any other function. In this case, Println
  can call it.

* EXERCISE (errors) - implemented based on the instructions. The error type does not have to be a struct. In
  this case, it is a float64. Similar to the previous lesson trying to emphasize golang best practices around
  error generation and handling.

* readers - the io.Reader interface has many implementations offered from golang. In this sample, there was no
  need to build an implementation.

* EXERCISE (readers) - a Reader implementation was the exercise.

* EXERCISE (readersrot13) - another Reader implementation - pretty fun to figure out. Maybe not the most efficient
  code, but a quick implementation.

* images - introduction to the Image interface.

* EXERCISE (images) - even with an empty struct as the type, the method implementations satisfy the exercise
  requirements. If you want to see the resulting image, the code has to be run in the browser version of
  the Tour of Go.

* goroutines - introduction to Go routines. Each routine is a lightweight thread accessing resources potentially
  shared by other routines.

* channels - these are like Java streams. You send and receive values with the operator, <-. You have to use
  make() like with maps and slices. Sends and receives block until the other side is ready, and that makes it very
  useful for routine synchronization without explicit locks. So the sample distributes work among 2 routines.
  When each routine is done, it sends an int on a channel. Both x and y ints are waiting to receive from the
  channel. 2 senders, 2 receivers - remember the blocking. So I start 2 threads to do work and then I wait in the
  caller until both threads are done. AND I have to receive results from both threads. Do the same in Java
  and compare the code :).

* channels (buf) - you can pass in a 2nd argument to make() which specifies a buffer length. When you do this,
  golang uses a buffer with the channel. Sends block until the buffer is full. Receives block if the buffer
  is empty. So try a couple of code mods and see what happens. First, try sending a 3rd int on the channel before
  the first Println(). Back to the original code (without the 3rd int send), try adding a 3rd Println(). The
  advantage of a buffer is to allow heavy traffic even though a receiver might not be ready. Without a buffer,
  the channel blocks on just a single send. With a buffer, sends can pile up in a buffer and receivers can pull
  when they are ready.

* channels (close) - back to the keyword range used in a loop. It works on a channel, but the channel has to
  be closed for the loop to end. Channels don't have to be closed unless a receiver needs to be told. And
  it is a good practice to only have a sender close a channel. Sending on a closed channel causes a runtime
  panic. The sample creates a buffered channel (and the cap() function can be used on a buffered channel).
  A routine is created and then main() waits on the range keyword. When the buffer has 10 ints, the loop
  activates because range can receive. Just before the routine shuts down, it closes the channel. This tells
  range to kick out of the loop. Cool. Again, write the same functionality in Java and compare the code :).

* channels (select) - another way to do the previous lesson without using close(). select waits until a
  channel operation is ready, and if multiple channel operations are ready the channel operation order
  is not guaranteed. Create 2 channels - not buffered. One channel is the main processing channel and the other
  is a "quit" channel. fibonacci() has an endless loop sending ints on the main processing channel
  and receiving ints on the "quit" channel. When one "quit" int is received, fibonacci() ends which also
  ends main(). The Go routine receives 10 ints from the main processing channel and sends one int on the
  "quit" channel.

* channels (selectdef) - just showing that a default can be used in a select. This will activate when all
  channel operations on the select are blocked.

* EXERCISE (concur) - main exercise for the Concurrency section of the tour. Maybe not the best code, but
  it was easy to implement and satisfies requirements.

* mutex - so channel communication is great for synchronizing routine execution. But if you don't need
  execution synchronization and only need data synchronization, golang offers a mutex. The sample
  shows this by starting 1000 routines all interested in the same data.

* EXERCISE (mutex) - last exercise building on previous lesson.

## keys-remote.sh

In a DEV container separate from the container used by VS Code, use [keys-remote.sh](https://github.com/kengibson1111/tour-of-go/blob/master/keys-remote.sh) to validate remote Git connectivity. If results look good, then any `git remote -v` error from VS Code is not related to how Git is interacting with Windows 11 Pro.

## Summary

The tour is pretty fun. Hopefully it will help others get off the ground and start exploring Go.

***

[^1]: Icon provided by [Icons8](https://icons8.com/).
