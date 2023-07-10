# moretypes - slices (literal)

This shows how a slice literal is very close to an array literal. In a single statement, the array literal is first created with a fixed size to match the array literal values. Then a slice is also created that has a pointer to the array literal along with a low-bound index and a high-bound index that represents all of the values in the array literal.
