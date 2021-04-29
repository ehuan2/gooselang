## Gooselang
This is Gooselang! An esoteric functional (complete) language with lazy evaluation. It supports both anonymous functions and global functions. All functions must take in a single parameter, ie it's curried.

Here are the commands:
| Syntax | Description | usage |
| --- | --- | --- |
| Honk | begins a body declaration, equivalent to ( | Honk ... honK |
| honK | ends a body declaration equivalent to ) | Honk ... honK |
| Gosling | anonymous function declaration, same as lambda, followed by variable name | Fowl var-name Honk ... honK |
| Goose | global statement, always followed by a var name or a Fowl | Goose var-name or Goose Fowl var-name Honk ... honK |
| HONK | applies an argument to a function | fn HONK arg |
| Hatch | turns some specially structured functions into a Go number and prints it | see more about lambda calculus' representation of binary representation of numbers |
| Lay | turns a Go natural number into a specially structured function of lambda calculus' binary representation | see more about lambda calculus |
| FLY | exits the program immediately | FLY |

It runs through the entire file, printing each structure that is not within a Goose structure. Note that Goose's must be declared before use.
Also, HONK is right applied first, so x HONK y HONK z => (x (y z))

Usage:
```
gooselang <filename>
gooselang
```

gooselang opens up the repl.

Formal syntax:
```
program = stmts-or-goose...
stmt = String | Gosling String Honk stmt honK | stmt HONK stmt | FLY
goose = Goose String stmt
```