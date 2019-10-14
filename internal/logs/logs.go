package logs

import (
	"log"
	"os"

	"github.com/efreitasn/customo"
)

// Error is the logger used for printing errors.
var Error = log.New(os.Stderr, customo.Format(
	"ERR: ",
	customo.AttrBold,
	customo.AttrFgColor4BitsRed,
), 0)

// Success is the logger used for printing success messages.
var Success = log.New(os.Stdout, customo.Format(
	"EVT: ",
	customo.AttrBold,
	customo.AttrFgColor4BitsGreen,
), 0)
