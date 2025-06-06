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
