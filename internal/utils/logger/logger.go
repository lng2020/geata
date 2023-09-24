// Copyright 2023 lng2020. All rights reserved.
// SPDX-License-Identifier: MIT

package logger

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func Init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}
