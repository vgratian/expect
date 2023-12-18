package internal

import (
	"fmt"
	"strings"
	"testing"
)

const (
	colorGreen = `[32m`
	colorRed   = `[31m`
	colorClear = `[0m`
	colorBold  = `[1m`
	statusPass = colorGreen + colorBold + "PASS" + colorClear
	statusFail = colorRed + colorBold + "FAIL" + colorClear
)

func prefixFromLabels(labels ...string) string {
	if len(labels) > 0 {
		return strings.Join(labels, ", ")
	}
	return ""
}

func formatMessage(status, prefix, format string, values ...interface{}) string {
	if prefix != "" && format != "" {
		prefix = prefix + ": "
	}
	if format != "" {
		return status + " - " + prefix + fmt.Sprintf(format, values...)
		//return fmt.Sprintf("%s - %s: %s", status, prefix, fmt.Sprintf(format, values...))
	}
	//if format == "" {
	//	return fmt.Sprintf("%s - %s", status, prefix)
	//}
	//return fmt.Sprintf("%s - %s: %s", status, prefix, fmt.Sprintf(format, values...))
	return status + " - " + prefix
}

func logPass(t *testing.T, format string, labels []string, values ...interface{}) {
	t.Log(formatMessage(statusPass, prefixFromLabels(labels...), format, values...))
}

func logFail(t *testing.T, format string, labels []string, values ...interface{}) {
	t.Error(formatMessage(statusFail, prefixFromLabels(labels...), format, values...))
}
