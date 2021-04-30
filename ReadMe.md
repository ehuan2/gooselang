## Gooselang
This is Gooselang! An esoteric functional (complete) language with lazy evaluation. It supports both anonymous functions and global functions. All functions must take in a single parameter, ie it's curried.

Here are the commands:
| Syntax | Description | usage |
| --- | --- | --- |
| Honk | begins a body declaration, equivalent to ( | Honk ... honK |
| honK | ends a body declaration equivalent to ) | Honk ... honK |
| HoNk | anonymous function declaration, same as lambda, followed by variable name | Fowl var-name Honk ... honK |
| HoNK | global statement, always followed by a var name or a Fowl | Goose var-name or Goose Fowl var-name Honk ... honK |
| HONK | applies an argument to a function | fn HONK arg |
| FLY | exits the program immediately | FLY |

It runs through the entire file, printing each structure that is not within a Goose structure. Note that Goose's must be declared before use.
Also, HONK is reversed, ie if it's f(a), we write a HONK f, f(a)(b), b HONK a HONK f

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