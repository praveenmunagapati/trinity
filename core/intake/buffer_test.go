package intake

import "testing"

func TestBuffer(t *testing.T) {
	buffer := BufferPool{}
	buffer.WriteString("something")
	buffer.WriteToFile("buffer_test.txt")
}
