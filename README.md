# go-lang-tutorial

Go language tutorial

```go
package main

import "fmt"

func main() {
   /* This is my first sample program. */
   fmt.Println("Hello, World!")
}
```

The example above shows a simple go program.
The statement `package main` declares the program that we're writing as main package.
The statement `import "fmt" ` imports a package called fmt. fmt is an inbuilt package of go language.
The line `fmt.Println("Hello, World!")` uses the Println method of fmt package to print the string "Hello world" to the console.

## Variables

Variables in go language are declared using the below syntax. The below example shows the declaration of variable of string type.

```go
 var name string ="John"
```

We declare the variable using the `var` keyword followed by the name of the variabe (in this case `name`) and type of the variable (In the above case it's a `string`)

Similarly we can delcare a variable of interger type as below

```go
 var myAge int= 36
```

### Other integer variable types

```go
var myVar1 int8= 12         // 8 bit integer
var myVar2 int16= 345       // 16 bit integer
var myVar3 int32= 374384    // 32 bit integer
var myVar4 int64= 34399343  // 64 bit integer
var myVar5 uint= 32         // unsigned integer
var myVar6 uint8=87         // 8 bit unsigned integer
var myVar7 uint16=456       // 16 bit unsigned integer
var myVar8 uint32=389393    // 32 bit unsigned integer
var myVar9 uint64= 390893934  // 64 bit unsigned integer
```

### Floating point variables

Floating point varibles are declared using the keywords `float32` and `float64`

```go
var pi float32 = 3.14159
var myFloatVar float64= 4.4838894858530
```

### Boolean type

Boolean type variable are declared using the keyword `bool` as shown in the below snippet

```go
var myBooleanVal bool= true
var isUserLoggedIn bool= false
```

## Declare multiple variables in a single statement

We can declare multiple variables in a single statement using the comma separated notation as shown below. Here the variables a,b,c and d are assigned the values 1,3,5 and 7 respectively.

```go
  var a, b, c, d int = 1, 3, 5, 7
```

Multiple varible declarations and their type can declared as shown below for greater readability.

```go
var (
     a int
     b int = 1
     c string = "hello"
   )
```

### Type inference

We can declare variables without explicitly using a type keyword. For example the following snippet is perfectly valid.

```go
  var a, b = 6, "Hello"
```

Here the variable `a` is assigned the value of `6` and the variable `b` is assigned the string value of `"Hello"`. In this case the type of the variable is inferred by the go compiler based on the value we assign.

### constants

If a variable is readonly and doesn't need to change its value we can declare the variable as a `const`.

```go
const PI float32= 3.14
```

## Arrays

Arrays are used to store the multiple values of the same type. In go language an array can be declared as follows.

`var array_name= [array_length]data_type{array_values}`

We can also make the go compiler infer the length of the array from the values.

`var array_name = [...]datatype{values} // here length is inferred`

Example:

```go
var myArr1= [3]int{1,2,3}  // array of length 3 with 3 values

var myArr2= [...]int{5,4,3,2}   // Here the length of the array is infered from the values

var fruits= [3]string{"Apple","Banana","Mango"}  // Array of string data type
```

### Access array elements

In Go language arrays start at zero index. An element of an array can be accessed using the below syntax.

`array_name[numerical_index]`

For example:

```go
package main
import ("fmt")

func main() {
var fruits= [3]string{"Apple","Banana","Mango"}

  fmt.Println(fruits[0])   // prints Apple
  fmt.Println(fruits[2])   // prints Mango
}
```

We can also choose to initialize the array at specific indexes only.

for example, in the below snippet an array of length 5 is declared but the values are not assigned. So by default 0 value is assigned for each index.

```go
  var arr1 = [5]int{} // not initializeds
```

### Partially initialized array

Whereas in the below code, the array is partially initialized. Even though the length of the array is declared as 5, only 0 and 1 index values are initialized.

```go
  var arr2 = [5]int{1,2} //partially initialized
```

### Initialize only specific elements

In go language it's possible to initialize only the specific elements of the array.

For example in the below code we assign the 2nd element to 10 and 4th element to 40

```go
  arr1 := [5]int{1:10,3:40}

```

### Length of an array

To find the length of an array we can use the `len()` method.
For example:

```go
  package main
import ("fmt")

func main() {
  var arr1 = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

  fmt.Println(len(arr1))
}

```

# Slices

Slices are similar to arrays but more more flexible and dynamic. Unlike arrays slices can grow and shrink in size.

The syntax to create a slice is

`slice_name :=[]datatype{values}`

For example, the below code shows the code for a slice of length 0 and capacity 0

```go
myslice := []int{}
```

The below code shows the syntax for a slice of length 3 and capacity 3.

```go
myslice := []int{1,2,3}
```

We can use the `len()` and `capacity()` methods to get the length and capacity of the slice.

For example:

```go
  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
```

Here are the differences between arrays and slices

| Feature           | Arrays                            | Slices                            |
| ----------------- | --------------------------------- | --------------------------------- |
| Declaration       | `var a [5]int`                    | `var s []int`                     |
| Length            | Fixed at compile time             | Dynamic, can grow or shrink       |
| Memory allocation | Entire array stored directly      | Points to an array under the hood |
| Type              | `[5]int` (length is part of type) | `[]int` (no fixed length)         |
| Example           | `[3]int{1, 2, 3}`                 | `[]int{1, 2, 3}`                  |

### Create a slice from an array

Since slices use arrays under the hood, we can create a slice from an array as shown below.

```go

   myFruits :=[5]={"Apple","Banana","Grapes", "Mango", "Strawberry"}
   mySlice := myFruits[1:4]  // selects [Banana Grapes Mango]

```

### Create slice using the make method

A slice can also be created using the make method.

`slice_name := make([]type,length, capacity)`

If the capacity property isn't provided, it will be equal to the length.
For example, the below code makes a slice of length 5 and capacity 10.

```go

mySlice :=make([]int,5,10)
```
