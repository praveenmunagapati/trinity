package intake

import (
	"bytes"
	"io/ioutil"
)

/*
BufferPool is the storage for a pool of bytes
being read from a file. IntakeBuffers are recycled
during parse loops to avoid allocation cost
*/
type BufferPool struct {
	bytes.Buffer
}

func (buf *BufferPool) WriteToFile(path string) {
	ioutil.WriteFile(path, buf.Bytes(), 0644)
}
