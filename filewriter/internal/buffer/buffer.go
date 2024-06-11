package buffer

import (
	"bytes"
	"sync"
)

type Buffer struct {
	bytes.Buffer
}

var buffers = sync.Pool{
	New: func() interface{} {
		return new(Buffer)
	},
}

// GetBuffer returns a new, ready-to-use buffer.
func GetBuffer() *Buffer {
	b := buffers.Get().(*Buffer)
	b.Reset()
	return b
}

// PutBuffer returns a buffer to the free list.
func PutBuffer(b *Buffer) {
	if b.Len() >= 256 {
		// Let big buffers die a natural death, without relying on
		// sync.Pool behavior. The documentation implies that items may
		// get deallocated while stored there ("If the Pool holds the
		// only reference when this [= be removed automatically]
		// happens, the item might be deallocated."), but
		// https://github.com/golang/go/issues/23199 leans more towards
		// having such a size limit.
		return
	}

	buffers.Put(b)
}
