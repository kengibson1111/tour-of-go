# tour-of-go

golang.org offers a Tour of Go. I took the tour on my laptop. Not the cleanest exercise
code, but I was just trying to satisfy the exercise requirements.

Lessons

* packages - illustrates a factored import and how exported functions from packages have
  to be capitalized. First look at Println. mains run from package main.

* imports - emphasizes factored import. First look at formatted printing.

* exports - shows how you can only access exported functions in a package. The exercise
  starts with an attempt to use math.pi instead of math.Pi.

* functions - shows the use of a private function.

* functions (params) - same as the previous lesson, but shows type sharing.

* functions (multiple) - this shows how a function can return any number of values.

* functions (named) - return values can be named in the function declaration. Then the
  return values can be referenced and assigned in the function. Pretty cool.

* variables - this shows variable declarations at a package and function scope.

* variables (init) - demonstrates initializers with variable declarations. The declaration at
  the function scope omits the type. And in that case, the variable infers the type of the
  initializer.

* variables (short) - introduces the short assignment (:=) instead of a variable declaration
  with an initializer omitting the type. For variables, this is only allowed at the function
  scope.

* types (basic) - this shows commonly used types. A factored var is used. First use of const
  with type inference. Notice const does not use short assignment like a variable. = instead
  of :=. Although not stated yet in the lessons, uppercase constants and variables at package
  scope could be exported from a package just like a function. Not sure yet.

* zero - if a var is not initialized with a value, it is assigned zero. 0 for numerics, false
  for bools, "" for strings.

* types (conversion) - example showing the required explicit type conversion. If you remove
  the float64 or uint type conversion in the sample, you'll get a compile error.

* types (inference) - simple example showing type inference. Change the variable value and see
  how the type changes.

* constants - even though const has been seen before, this is the official lesson. Still not sure
  about package-level uppercase constants and variables being exported from the package.

* constants (numeric) - uses a factored const. Shows bit shifting and type inference, but the type
  conversion is inconsistent. Just looking at the code, I would assume Big is a float. And because
  Small uses Big in its assignment, I would also assume Small is a float. In the types (conversion)
  lesson, the conversion between types had to be explicit. In this lesson, Small is converted to an
  int and I don't get compile or runtime errors.

* forloop - basic for loop that looks like Java except for the variable short assignment :=.

* forloop (while) - because init and post statements are optional, this is how to make a for loop
  act like a while.

* whileloop - same as previous sample only drop the leading and trailing semicolons.

* forever - building on the previous sample, omit the condition and you have the infinite for loop.

* ifstmt - basics of the if statement.

* if (short) - the if statement can have a "short assignment" init before the condition. Like a for loop
  without a post statement.

* if (else) - adding the else to the previous sample.

* EXERCISE (loops) - not the cleanest code, but it did meet the exercise requirements.

* switch - basic switch statement.

* switch (order) - this emphasizes switch statement order and the built-in case break when a statement
  evaluates to true. In Java, you can stack case statements - not so with golang.

* switch (nocond) - you can control the condition statements at the case level instead of the top-level switch.
  Slick way of building long if-then-else chains.

* defer - you can defer code to run just after the function is popped off the call stack. Handy for resource
  cleanup.

* defer (stack) - this shows that deferred code has to play by call stack rules. Makes sense.

* pointers - oh yeah, golang has pointers. This is where I think golang starts to kick ass in the runtime. The
  pointers are not quite like C pointers because there is no arithmetic. But the efficiencies are there, and
  pointer arithmetic can venture into illegal process memory space anyway :). Function call stacks can be really
  tight like in C. This lesson is focusing only on a pointer to a primitive type.

* structs - and the fun continues. objects? Uh, no. And there is a way to define and export a "method" meant
  specifically for a struct, but that's later. Functions, pointers, structs - back to the future. For this lesson
  I would assume that the struct instance is being created on the stack (not heap). Whether or not
  fmt.Println copies the struct instance or just uses a pointer to the struct for the call stack is unknown
  at this point.

* structs (fields) - fields are accessed via dot notation.

* pointers (structs) - this is showing that a pointer to a struct can be used to make changes in the struct
  instance. C makes a field access distinction between a struct and a pointer. golang doesn't do that - dot
  notation all the way. In this sample, it looks like a a struct instance is created on the stack and the
  pointer value is stack address space.

* structs (literals) - a few examples of creating a struct instance using literals. In the pointer case, the
  struct instance is created first and then the pointer is assigned. And I would assume for the Println
  call at runtime, 3 struct instances are being copied from the stack to the function call stack and 1 pointer
  value is being copied from the stack to the function call stack. Doesn't seem like a big deal when
  a struct is so small, but it could be a big deal with more complex structs.

* arrays - simple example creating an array on the stack. arrays have a hard limit - you specify the length
  and the array is only that length. Thoughts on the golang string implementation: just an assumption, but it
  looks like a string is a pointer to a char array. Memory for the char array is not initialized until the string
  is initialized with a value. Then when a string is initialized (like in the sample code), char array memory
  is allocated and set. So the sample code creates an array of 2 char array pointers on the stack. Then golang
  allocates memory for the char arrays on the stack or heap depending on the underlying implementation. This
  seems like a small point, but when you assign and reassign string (char array pointer) values this could be
  expensive for the runtime. golang hides all of the string complexity in the language and runtime, but the
  complexity has to be addressed eventually.

* slices - this shows how effective pointers are since a slice is a pointer to an array. Building on comments
  from the previous lesson, I believe the golang implementation of a string is a slice pointing to a char array.
  The sample code creates an array of ints using literals. Then the pointer to the int array is created. That is
  the slice. I would think if a function needs an array as a parameter, you want to use a slice for that because
  that is the best optimization of the function call stack.

* slices (of slice) - and since a slice is a pointer, the array type can be anything including a slice which,
  again, is a pointer. Similar to pointer of pointers in C. There is a lot going on in the sample even though
  much is hidden in the language. Each literal string is a slice to a char array. So the memory for the
  char array slices (strings) have to be created first. Then each string array is created followed by a slice
  for each string array. 9 char arrays, 9 slices to char arrays (strings), 3 string arrays, 3 slices to string
  arrays. So far. Then the array of string slices is created followed by the slice for that array. And for
  the underlying slice implementation, still not sure if the memory is on the stack or heap based on the sample
  code.

* slices (range) - when a slice subset is specified using [], a new pointer to the same array is created. This
  is where golang has made a slice more than just a pointer. Only the specified range is visible in the newly
  created slice.
