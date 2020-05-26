package redis

import (
	"bufio"
	"io"
	"strconv"
)

// RESP
// Simple strings are used for common server replies such as “OK” (after a successful write command) or “PONG” (the successful response to the PING command).
// Bulk strings are returned for almost all single-value read commands such as GET, LPOP, and HGET. Bulk strings are different from simple strings in that they can contain anything — newlines, control characters, or even valid RESP, all without being escaped or encoded.
// Integers are used as the reply for any kind of counting command such as STRLEN, HLEN, or BITCOUNT.
// Arrays can contain any number of RESP objects, including other arrays. This is used to send commands to the server, as well as for any reply that returns more than one element such as HGETALL, LRANGE, or MGET.
// Errors are returned whenever Redis encounters an error while handling your command, such as when trying to run a command against a key holding the wrong type of data.

// + is the simple string prefix
// Errors are prefixed with -
// Integers are prefixed with :
// $ is the bulk string prefix
// \r\n is the line terminator

// Suppose a simple client only needs to write an array of bulk strings for sending commands to Redis
var (
	arrayPrefixSlice = []byte{'*'}
	bulkStringSlice  = []byte{'$'}
	lineEndingSlice  = []byte{'\r', '\n'}
)

type RESPWriter struct {
	*bufio.Writer
}

func NewRESPWriter(writer io.Writer) *RESPWriter {
	return &RESPWriter{
		Writer: bufio.NewWriter(writer),
	}
}

func (w *RESPWriter) WriteCommand(args ...string) error {
	//head
	w.Write(arrayPrefixSlice)
	w.WriteString(strconv.Itoa(len(args)))
	w.Write(lineEndingSlice)
	//*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n
	for _, arg := range args {
		w.Write(bulkStringSlice)
		w.WriteString(strconv.Itoa(len(arg)))
		w.Write(lineEndingSlice)
		w.WriteString(arg)
		w.Write(lineEndingSlice)
	}

	return w.Flush()
}
