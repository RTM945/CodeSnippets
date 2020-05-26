package redis

const (
	SimpleString = '+'
	BulkString   = '$'
	Integer      = ':'
	Array        = '*'
	Error        = '-'
)

// var (
// 	ErrInvalidSyntax = errors.New("resp: invalid syntax")
// )

// type RESPReader struct {
// 	*bufio.Reader
// }

// func NewRESPReader(reader io.Reader) *RESPReader {
// 	return &RESPReader{
// 		Reader: bufio.NewReaderSize(reader, 32*1024),
// 	}
// }

// func (r *RESPReader) readLine() (line []byte, err error) {
// 	line, err = r.ReadBytes('\n')
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(line) > 1 && line[len(line)-1] == '\r' {
// 		return line, nil
// 	}
// 	return nil, ErrInvalidSyntax
// }

// func (r *RESPReader) getCount(line []byte) (int, error) {
// 	end := bytes.IndexByte(line, '\r')
// 	return strconv.Atoi(string(line[1:end]))
// }

// func (r *RESPReader) readBulkString(line []byte) ([]byte, error) {
// 	count, err := r.getCount(line)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if count == -1 {
// 		return line, nil
// 	}
// 	buf := make([]byte, len(line)+count+2)
// 	copy(buf, line)
// 	_, err = io.ReadFull(r, buf[len(line):])
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf, nil
// }

// func (r *RESPReader) readArray(line []byte) ([]byte, error) {
// 	count, err := r.getCount(line)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i := 0; i < count; i++ {
// 		buf, err := r.ReadObject()
// 		if err != nil {
// 			return nil, err
// 		}
// 		line = append(line, buf...)
// 	}
// 	return line, nil
// }

// func (r *RESPReader) ReadObject() ([]byte, error) {
// 	line, err := r.readLine()
// 	if err != nil {
// 		return nil, err
// 	}
// 	switch line[0] {
// 	case SimpleString, Integer, Error:
// 		return line, nil
// 	case BulkString:
// 		return r.readBulkString(line)
// 	case Array:
// 		return r.readArray(line)
// 	default:
// 		return nil, ErrInvalidSyntax
// 	}

// }
