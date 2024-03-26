package infrastructures

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strings"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func VerbTitle(message string) {
	maxChar := 50
	messageLength := len(message) + 2
	maxLineChar := maxChar - messageLength
	modLineChar := int(math.Mod(float64(maxLineChar), 2))
	leftLineChar := (maxLineChar - modLineChar) / 2
	rightLineChar := maxChar - (leftLineChar + messageLength)
	Factory.Log.Warn(strings.Repeat("=", leftLineChar) + " " + message + " " + strings.Repeat("=", rightLineChar))
}

func VerbTitleF(message string) {
	VerbTitle(message)
}

func VInfo(message string) {
	Factory.Log.Info(message)
}

func VWarn(message string) {
	Factory.Log.Warn(message)
}

func VErr(err error) {
	errorLogStr := getStackTraceStr(err)
	Factory.Log.ErrorF(map[string]interface{}{"stacktrace": errorLogStr}, err.Error())
}

func VErrStr(stackTrace string, err string) {
	Factory.Log.ErrorF(map[string]interface{}{"stacktrace": stackTrace}, err)
}

func getStackTraceStr(err error) string {
	errResult := errors.WithStack(err)
	stackTraces := errResult.(stackTracer).StackTrace()
	max := len(stackTraces)
	if max >= 7 {
		max = 7
	}

	errorLogStr := ""
	tracesCount := len(stackTraces)

	for i := 0; i < len(stackTraces); i++ {
		f := stackTraces[tracesCount-i-1]
		errorLogStr += fmt.Sprintf("%+s:%d\n", f, f)
	}
	return errorLogStr
}
