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
