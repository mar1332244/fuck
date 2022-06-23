package main

import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

var (
    cellArr [CELL_COUNT]byte
    cellPtr int
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("argument error:\n\tno file given")
        os.Exit(ARGUMENT_ERROR)
    }
    if !strings.HasSuffix(os.Args[1], ".bf") {
        fmt.Println("argument error:\n\tunsupported file type")
        os.Exit(ARGUMENT_ERROR)
    }
    interpretSrc(readSrc(os.Args[1]))
}

func readSrc(fname string) string {
    var src strings.Builder
    fp, err := os.Open(fname)
    if err != nil {
        fmt.Printf("io error:\n\t%s\n", err.Error())
        os.Exit(IO_ERROR)
    }
    defer fp.Close()
    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        for _, char := range scanner.Text() {
            if !strings.ContainsRune("+-<>.,[]", char) {
                continue
            }
            src.WriteRune(char)
        }
    }
    if err = scanner.Err(); err != nil {
        fmt.Printf("io error:\n\t%s\n", err.Error())
        os.Exit(IO_ERROR)
    }
    fmt.Println(src.String())
    return src.String()
}

func interpretSrc(src string) {
    var i int
    var err error
    var c byte
    reader := bufio.NewReader(os.Stdin)
    jumpStack := new(Stack)

    for i = 0; i < len(src); i++ {
        switch src[i] {

        case CELL_PLUS:
            cellArr[cellPtr]++

        case CELL_MINUS:
            cellArr[cellPtr]--

        case MOVE_RIGHT:
            cellPtr++
            if cellPtr == CELL_COUNT {
                cellPtr = 0
            }

        case MOVE_LEFT:
            cellPtr--
            if cellPtr == -1 {
                cellPtr = CELL_COUNT - 1
            }

        case WRITE_BYTE:
            fmt.Printf("%c", cellArr[cellPtr])

        case READ_BYTE:
            c, err = reader.ReadByte()
            if err != nil {
                fmt.Printf("io error:\n\t%s\n", err.Error())
                os.Exit(IO_ERROR)
            }
            cellArr[cellPtr] = c

        case BEGIN_LOOP:
            jumpStack.Push(i)

        case END_LOOP:
            if jumpStack.Size == 0 {
                fmt.Println("syntax error:\n\tunmatched ]")
                os.Exit(SYNTAX_ERROR)
            }
            if cellArr[cellPtr] == 0 {
                jumpStack.Pop()
            } else {
                i = jumpStack.Peek()
            }
        }
    }

    if jumpStack.Size > 0 {
        fmt.Println("syntax error:\n\tunmatched [")
        os.Exit(SYNTAX_ERROR)
    }
}

func setUpRepl() {
}

func compileSrc(src string) {
}
