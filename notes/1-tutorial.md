# Chapter 1: Tutorial

1. Go is a compiled language. ```go run``` compiles source code from
one or more source files whose names end in .go, links it with
libraries, then runs the resulting executable file.

2. Go handles Unicode natively, so it can process text in all the world
languages.

3. Common commands, To:
 - save the build file for later use: ```go build file_name.go``` (thi creates an executable binary file, which can be executed without furthur processing)
 - simply compile and run: ```go run file_name.go```

4. Packages in golang
They are similar to modules and libraries in other languages.
A package has one or more .go source files in a single directory.
Each source file begins with a package declaration.
In our file, we did: ```package main``` - this tells which package
the file belongs to.
After this, there is a list of packages to be imported.

__fmt__ is the format package - prints formatted output and scanning
input. Println is a basic output function in fmt.

5. Package main
is special. Defines a standalone executable program, not a library.
within package main, function main is also special - it's where
execution of program begins - like in C. Program does whatever main
does, ONLY.

6. import safety restriction
"You must import exactly the packages you need. A program will not
compile if there are missing imports or if there are unnecessary ones. 
This strict requirement prevents references to unused packages from
accumulating as programs evolve."

7. Basic structure
```
package main                 // package declaration
import "whatever_you_want"   // import declarations
func main() {                // functions, variables
    ...
}
```
import declarations have to follow package declaration

8. (Imp.)
"Go does not require semicolons at the ends of statements or
declarations, except where two or more appear on the same line.
In effect, newlines following certain tokens are converted into
semicolons, so where newlines are placed matters to proper parsing
of Go code.
For instance,
the opening brace { of the function must be on the same line as the
end of the func declaration, not on a line by itself, and in the
expression x + y, a newline is permitted after but not before
the + operator.
Declaring a standard format by fiat eliminates a lot of pointless
debate about trivia and, more importantly, enables a variety of
automated source code transformation 
that would be infeasible if arbitrary formatting were allowed."

7. For loop:
```go
for initialization (optional); condition; post {
    // zero or more statements
}
```
around the 3 components of the loop - parenthesis not used - braces are
mandatory, and opening brace must be in the same line as the post statement.

 - Optional initialization statement is executed before the loop starts.
 - condition is a boolean expression, if evaluated true, the loop executes.
 - The post statement is executed after the body of loop executes - then condition is evaluated again and so on...
 - loop ends when the condition is false.

If there is no initialization and no post, semi-colon may be omitted entirely.
```go
for condition {
    //
}
```

This is an infinite loop in go
```go
for {
    //
}
// although the body of the loop may have logic to terminate the loop - break or return.
```

The blank
identifier may be used whenever syntax requires a variable name but program logic does not,
for instance to discard an unwanted loop index when we require only the element value. Most
Go programmers would likely use range and _ to write the echo program as above, since the
indexing over os.Args is implicit, not explicit, and thus easier to get right

8. Ways to declare a variable
```go
s := ""
var s string
var s = ""
var s string = ""
```
In practice, you should generally
use one of the first two forms, with explicit initialization to say that the initial value is
important and implicit initialization to say that the initial value doesn’t matter.


9. Every variable in Go is automatically initialized to a default state called its zero value
when declared without an initial value. For reference types, pointers, and interfaces, that
zero value is nil. For basic types, it is a non-nil default (like 0 or "").


10. Array vs Slice in GOLANG
**Array**                                       **Slice**
 - size is explicit: [5]int                     size is blank: []int
 - size is fixed at compile time; cannot grow   Dynamic; grows automatically when element added


11. os package in GO
os.Args is a string slice.
os.Args[0] is the name of the program
os.Args[1 onwards] are the words written in the command line.
```bash
$ go run file_name.go hello world
```
Here os.Args[1] = hello; os.Args[2] = world and so on if more cmd arguments!

```bash
apvshuk@computer:~/Documents/go-knd/exercises$ go run 1.2.go hello world hey world
0 /tmp/go-build3581099105/b001/exe/1.2
1 hello
2 world
3 hey
4 world
```
This is the output of exercise 1.2; which was to print index and value of each element of
os.Args.
Notice that the first element is the location/name of the program running; ensuing elements
are the command line arguments; one word is each element.




