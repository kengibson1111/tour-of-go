# methods - functions (pointers)

Just to drive home the concept, methods are functions. So the same arguments apply to passing a value of a type instance to prevent updates and a pointer of the type instance to allow updates.

In **methods - methods (pointers)**, if *Vertex was changed in Vertex in Scale()'s type receiver, dot notation hid a potential compile error. And, Scale() worked with a copy of a Vertex instance on the funtion call stack because the type receiver was a value. In this example, change the *Vertex to Vertex in Scale() and you should see a compile error. Fix it and try again.

It might be more advantageous to see the compile error first (like in this example) instead of waiting for a runtime error.
