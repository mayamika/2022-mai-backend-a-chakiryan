package upload

import "io"

type Upload struct {
	File        io.Reader
	Filename    string
	Size        int64
	ContentType string
}
