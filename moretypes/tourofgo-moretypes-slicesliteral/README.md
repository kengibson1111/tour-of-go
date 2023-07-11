# moretypes - slices (literal)

This shows how a slice literal is very close to an array literal. In a single statement, the array literal is first created with a fixed size to match the array literal values. Then a slice is also created that has a pointer to the array literal along with a low-bound index of zero, a length that matches the array literal size, and a capacity that also matches the array literal size.
