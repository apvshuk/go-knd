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
