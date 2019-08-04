package paasio

import (
	"io"
	"sync"
)

type counter struct {
	mux    *sync.Mutex
	nbytes int64
	nops   int
}

func newCounter() counter {
	return counter{mux: new(sync.Mutex)}
}
func (c *counter) count(nbytes int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.nops++
	c.nbytes += int64(nbytes)
}
func (c *counter) stats() (nbytes int64, nops int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.nbytes, c.nops
}

type readCounter struct {
	r io.Reader
	counter
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.r.Read(p)
	rc.count(n)
	return
}
func (rc *readCounter) ReadCount() (nbytes int64, nops int) {
	return rc.stats()
}

type writeCounter struct {
	w io.Writer
	counter
}

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	wc.count(n)
	return
}
func (wc *writeCounter) WriteCount() (nbytes int64, nops int) {
	return wc.stats()
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

// NewReadCounter returns an implementation of ReadCounter.  Calls to
// r.Read() are not guaranteed to be synchronized.
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r, counter: newCounter()}
}

// NewWriteCounter returns an implementation of WriteCounter.  Calls to
// w.Write() are not guaranteed to be synchronized.
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w, counter: newCounter()}
}

// NewReadWriteCounter returns an implementation of ReadWriteCounter.
// Calls to rw.Write() and rw.Read() are not guaranteed to be synchronized.
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(rw), NewWriteCounter(rw)}
}
