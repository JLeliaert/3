package httpfs

// Utility functions on top of standard httpfs protocol

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
)

// create a file for writing, clobbers previous content if any.
func Create(URL string) (WriteCloseFlusher, error) {
	_ = Remove(URL)
	err := Touch(URL)
	if err != nil {
		return nil, err
	}
	return &bufWriter{bufio.NewWriterSize(&appendWriter{URL}, 4*1024*1024)}, nil
}

type WriteCloseFlusher interface {
	io.WriteCloser
	Flush() error
}

// open a file for reading
func Open(URL string) (io.ReadCloser, error) {
	data, err := Read(URL)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewReader(data)), nil
}

func Touch(URL string) error {
	return Append(URL, []byte{})
}

type bufWriter struct {
	buf *bufio.Writer
}

func (w *bufWriter) Write(p []byte) (int, error) { return w.buf.Write(p) }
func (w *bufWriter) Close() error                { return w.buf.Flush() }
func (w *bufWriter) Flush() error                { return w.buf.Flush() }

type appendWriter struct {
	URL string
}

// TODO: buffer heavily, Flush() on close
func (w *appendWriter) Write(p []byte) (int, error) {
	err := Append(w.URL, p)
	if err != nil {
		return 0, err // don't know how many bytes written
	}
	return len(p), nil
}

// TODO: flush
func (w *appendWriter) Close() error {
	return nil
}
