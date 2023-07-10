# moretypes - slices (reference)

This builds upon **moretypes - slices** and again shows that a slice is a reference struct to an array. It has a pointer to the array along with a low-bound index and a high-bound index related to the array. Remember, slices do not store data like an array. Another important point is: when a slice makes a change in an array element, all other slices referencing the same array can see the change.

The sample code creates an array of strings using literals. Then two slices to the string array are created. The second slice makes a change to its first element (which is the array's second element). The first slice can see the change, and the array itself can see the change.
