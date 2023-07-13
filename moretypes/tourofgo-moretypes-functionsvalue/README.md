# moretypes - functions (value)

Functions are first-class citizens in Go. And, that opens up the design possibilities. In the sample, compute() has a function parameter. All compute() knows is that the function will accept two values and compute() will pass 3 and 4 as the values to the function parameter. Now any function can be passed that meets the parameter spec and it will do something with the values 3 and 4. math.Pow meets the spec, but you can also define your own function that also meets the spec. compute() doesn't care.
