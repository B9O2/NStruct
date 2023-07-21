package FileQueue

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type FileQueue struct {
	data      map[int]string
	f         *os.File
	reader    *bufio.Reader
	delim     byte
	lineCount int
	index     int
	max       int
}

func (q *FileQueue) Pop() (string, error) {
	if q.index <= q.max {
		result := q.data[q.index]
		delete(q.data, q.index)
		q.index += 1
		return result, nil
	} else {
		counter := q.readMore()
		if counter > 0 {
			return q.Pop()
		} else {
			return "", errors.New("no more data")
		}
	}
}

func (q *FileQueue) PopN(n int) {

}

func (q *FileQueue) readMore() int {
	counter := 0
	for i := 0; i < q.lineCount; i++ {
		line, err := q.reader.ReadString(q.delim)
		switch err {
		case nil:
			q.max += 1
			q.data[q.max] = line
			counter += 1
		case io.EOF:
			q.max += 1
			q.data[q.max] = line
			counter += 1
			q.f.Close()
			break
		default:
			q.f.Close()
			break
		}
	}
	return counter
}

func NewQueue(filePath string, delim byte, size int) FileQueue {
	f, err := os.Open(filePath)
	if err != nil {
		//fmt.Println("read file '" + filePath + "' failed.")
		return FileQueue{}
	} else {
		q := FileQueue{
			index:     1,
			delim:     delim,
			lineCount: size,
			reader:    bufio.NewReader(f),
			max:       0,
			data:      map[int]string{},
		}
		return q
	}
}
