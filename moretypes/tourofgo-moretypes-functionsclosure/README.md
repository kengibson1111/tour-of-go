# moretypes - functions (closure)

Closures can reference variables outside of their local function scope. So in the sample, adder() is returning a closure that can mutate its own sum variable. Seems like a potential design and implementation for basic map reductions.
