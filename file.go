package file

import (
	"io"
)

type File interface {
	io.Reader
	io.Writer
	io.Closer
	io.Seeker
	Name() string
}
