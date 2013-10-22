// Package randomart generates OpenSSH-like visual hashes
// Based on https://pthree.org/2013/05/30/openssh-keys-and-the-drunken-bishop/
package randomart

import (
	"bytes"
	"io"
)

const SSHChars = " .o+=*BOX@%&#/^"

type bishop struct {
	// FIXME: title?
	board      [][]byte
	y, x       int
	ymax, xmax int
	chars      string
}

func NewSSH() *bishop {
	return New(9, 17, SSHChars)
}

func New(y, x int, chars string) *bishop {

	board := make([][]byte, y)
	for i := range board {
		board[i] = make([]byte, x)
	}

	return &bishop{
		board: board,
		y:     (y - 1) / 2,
		x:     (x - 1) / 2,
		ymax:  y,
		xmax:  x,
		chars: chars,
	}
}

func (b *bishop) Write(buf []byte) (int, error) {

	n := len(buf)

	m := moves{data: buf}

	for {
		r, err := m.next()
		if err == io.EOF {
			break
		}

		moveSouth, moveEast := (r >> 1), (r & 1)

		if moveSouth == 1 && b.y < (b.ymax-1) {
			b.y++
		} else if moveSouth == 0 && b.y > 0 {
			b.y--
		}

		if moveEast == 1 && b.x < (b.xmax-1) {
			b.x++
		} else if moveEast == 0 && b.x > 0 {
			b.x--
		}

		b.board[b.y][b.x]++
	}

	return n, nil
}

func (b *bishop) String() string {

	xstart := (b.xmax - 1) / 2
	ystart := (b.ymax - 1) / 2

	var buf bytes.Buffer

	buf.Write([]byte("+-----------------+\n"))

	for y := range b.board {
		buf.WriteByte('|')
		for x := range b.board[y] {
			count := b.board[y][x]
			var ch byte
			if int(count) < len(b.chars) {
				ch = b.chars[count]
			} else {
				ch = b.chars[len(b.chars)-1]
			}
			// mark start and end points
			if x == xstart && y == ystart {
				ch = 'S'
			} else if x == b.x && y == b.y {
				ch = 'E'
			}
			buf.WriteByte(ch)
		}
		buf.Write([]byte{'|', '\n'})
	}
	buf.Write([]byte("+-----------------+"))

	return buf.String()
}

// small type to make extracting bits easier

type moves struct {
	data  []byte
	b     byte
	count int
}

func (m *moves) next() (byte, error) {
	if len(m.data) == 0 && m.count == 0 {
		return 0, io.EOF
	}

	if m.count == 0 {
		// no bits are valid in 'b', so refill it
		m.b = m.data[0]
		m.count = 8
		m.data = m.data[1:]
	}

	// return two lowest bits
	r := (m.b & 0x3)

	// mark as used
	m.b >>= 2
	m.count -= 2

	return r, nil
}
