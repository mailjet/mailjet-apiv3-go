package mailjet

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// sortOrder defines the order of the result.
type sortOrder int

// These are the two possible order.
const (
	SortDesc = sortOrder(iota)
	SortAsc
)

var debugOut io.Writer = os.Stderr

// filter applies a filter with the defined key and value.
func filter(key, value string) RequestOptions {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = strings.Replace(q.Encode(), "%2B", "+", 1)
	}
}

// sort applies the sort filter to the request.
func sort(value string, order sortOrder) RequestOptions {
	if order == SortDesc {
		value = value + "+DESC"
	}
	return filter("Sort", value)
}

// SetDebugOutput sets the output destination for the debug.
func SetDebugOutput(w io.Writer) {
	debugOut = w
	log.SetOutput(w)
}
