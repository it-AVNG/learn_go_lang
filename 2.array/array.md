# Learn how GO's array works

## Init an array

It is simple by using ':=' operand, follow with `[length]type{value1,value2..}`.
We can also use `var` to declare the array.

- Partialy init an array with inferred length `[...]type{value1,value2,..}`
- Partially init an array with incomplete value, GO fills the rest with empty value.

## Access value and edit value

Array value can be access by normal indexing with `array[i]`.
To assign the array value, treat it like a variable.

Access value while initiating by the following syntax `array:=[length]type{index:value}`.
The array length can be checked by `len(array)`

# Learn how GO's slices works

Slices are array but more flexible.

## Init a slice:

- Init a slice : `slice_name := []type{value}`
- Init a slice from an Array: `slice_name := array[index_start:index_end]`
- Init a slice with `make()`: `slice_name := make([]type, length, capacity)`

As a norm, "slicing" an Array to create a Slice, does not include the value at `index_end`.
`capacity`, inferred that the underlying array of the slice is has a capacity to store more values.
For example: if a slice is set for length of 3 but underlying capacity is 5. When there are "overflow",
the memory allocation to the variable will not be adjusted, because it has been provided a `capacity`.

## Slice operation:

A Slice is an Array, any value can be address with indexing, except for out-of-bound.
To append a Slice we use `append()`, syntax are as follow:

```go
slice_name = append(slice_name, element1, element2, ...)
```

When `append()` with number of elements that is over the capacity of a slice, Go will automate allocate new capacity.
