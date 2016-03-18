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
  the function scope omits the type. And in that case, the variable assumes the type of the
  initializer.

* variables (short) - introduces the short assignment (:=) instead of a variable declaration
  with an initializer omitting the type. For variables, this is only allowed at the function
  scope.
