package main

const (
    EXIT_SUCESS int = iota
    ARGUMENT_ERROR  = iota
    IO_ERROR        = iota
    SYNTAX_ERROR    = iota
    EXIT_CODE_COUNT = iota
)

const (
    CELL_PLUS  byte = '+'
    CELL_MINUS byte = '-'
    MOVE_LEFT  byte = '<'
    MOVE_RIGHT byte = '>'
    BEGIN_LOOP byte = '['
    END_LOOP   byte = ']'
    READ_BYTE  byte = ','
    WRITE_BYTE byte = '.'
)

const (
    CELL_COUNT int = 30000
)
