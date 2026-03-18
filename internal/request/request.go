package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var ErrBadRequestLine = errors.New("malformed request line")
var crlf = []byte("\r\n")

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

func (rl *RequestLine) ValidHttp() bool {
	return rl.HttpVersion == "HTTP/1.1"
}

type Request struct {
	RequestLine RequestLine
}

func parseRequestLine(data *[]byte) (*RequestLine, error) {
	idx := bytes.Index(*data, crlf)
	if idx == -1 {
		return nil, ErrBadRequestLine
	}

	startLine := (*data)[:idx]

	*data = (*data)[idx+len(crlf):]

	parts := bytes.Fields(startLine)
	if len(parts) != 3 {
		return nil, ErrBadRequestLine
	}

	rl := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion:   string(parts[2]),
	}

	if !rl.ValidHttp() {
		return nil, ErrBadRequestLine
	}

	return rl, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("unable to io.ReadAll "), err)
	}

	rl, err := parseRequestLine(&data)
	if err != nil {
		return nil, err
	}

	return &Request{RequestLine: *rl}, nil
}
