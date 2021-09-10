package trie

import (
	"container/list"
	"fmt"
)

type Words struct {
	buf *list.List
	segWords <-chan string
	front, curr *list.Element
	res chan string
	exit bool
}

func NewWords(key string) *Words {
	ws := &Words{
		buf: list.New(),
		segWords: parse(key),
		res: make(chan string),
		exit: false,
	}
	go ws.waitReq()
	return ws
}

func (ws *Words) Iter() (s string, ok bool) {
	if ws.exit {
		return
	}
	ws.res <- ""
	s, ok = <-ws.res
	return
}

func (ws *Words) waitReq() {
	for range ws.res {
		// ws.Dump("iter")
		ws.front = ws.buf.Front()
		if ws.front == nil {
			w, ok := <-ws.segWords
			if !ok {
				break
			}
			ws.front = ws.buf.PushBack(w)
			ws.curr = nil
			// fmt.Printf("<- %s\n", w)
			ws.res <- w
			continue
		}

		if ws.curr == nil {
			w, ok := <-ws.segWords
			if !ok {
				if n := ws.front.Next(); n == nil {
					break
				}
				ws.res <- ""
				continue
			}
			ws.curr = ws.buf.PushBack(w)
		}

		w := ws.curr.Value.(string)
		ws.curr = ws.curr.Next()
		// ws.Dump("after iter")
		// fmt.Printf("<- %s\n", w)
		ws.res <- w
	}

	ws.exit = true
	close(ws.res)
}

func (ws *Words) Reset() {
	// ws.Dump("before reset")
	ws.front = ws.buf.Front()
	if ws.front != nil {
		ws.buf.Remove(ws.front)
		ws.front = ws.buf.Front()
		ws.curr = ws.front
	} else {
		ws.curr = nil
	}
	// ws.Dump("after reset")
}

func (ws *Words) Dump(prompt string) {
	fmt.Printf("%s: front", prompt)
	if ws.front == nil {
		fmt.Printf("[nil], ")
	} else {
		fmt.Printf("(%s), ", ws.front.Value.(string))
	}
	fmt.Printf("curr")
	if ws.curr == nil {
		fmt.Printf("[nil], ")
	} else {
		fmt.Printf("(%s), ", ws.curr.Value.(string))
	}
	fmt.Printf("\n")
}
