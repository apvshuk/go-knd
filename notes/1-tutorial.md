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