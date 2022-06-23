# fuck

A set of utilities for working with the Brainfuck programming language written in Go

## The fuck command

```
$ fuck
Welcome to fuck v1.0.1
Type .help for more info or .exit to exit
(fuck) >+++++++++[<++++++++++>-]<+++++++.
a
(fuck)
```

## Working with fuck as a library

```
package main

import (
    "github.com/mar1332244/fuck"
)

func main() {
    var program fuck.Array = fuck.NewArray(30000)
    program.WriteSnippet(">+++++++++[<++++++++++>-]<+++++++")
    program.WriteCommand('.')
    program.Execute() // a is printed
    program.SaveTo("a.bf")
    program.Compile("a")
}
```
