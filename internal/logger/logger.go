package logger

import (
    "log"
    "os"
)

var L = log.New(os.Stderr, "[guardian] ", log.LstdFlags)
