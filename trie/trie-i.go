package trie

import (
	gotrie "github.com/derekparker/trie"
	"strings"
	"bytes"
)

type Occurring struct {
	Count int `json:"count"`
	Meta interface{} `json:"meta"`
}

type Trie struct {
	t *gotrie.Trie
}

func NewTrie() *Trie {
	return &Trie{t:gotrie.New()}
}

func (t *Trie) Append(key string, meta interface{}) {
	t.t.Add(strings.ToLower(key), meta)
}

func (t *Trie) Remove(key string) {
	t.t.Remove(strings.ToLower(key))
}

func (t *Trie) Get(key string) (interface{}, bool) {
	node, ok := t.t.Find(strings.ToLower(key))
	if ok {
		return node.Meta(), true
	}
	return nil, false
}

func (t *Trie) PrefixSearch(key string) []string {
	return t.t.PrefixSearch(key)
}

func (t *Trie) Clear() {
	t.t = gotrie.New()
}

func (t *Trie) SimpleSearch(key string) map[string]*Occurring {
	res := make(map[string]*Occurring)

	buf := &bytes.Buffer{}

	for w := range parse(key) {
		lw := strings.ToLower(w)
		buf.WriteString(lw)
		prefix := buf.String()
		if t.t.HasKeysWithPrefix(prefix) {
			setMeta(t.t, res, prefix)
			continue
		}

		buf.Reset()
		buf.WriteString(lw)
	}

	return res
}

func (t *Trie) Search(key string) map[string]*Occurring {
	res := make(map[string]*Occurring)

	buf := &bytes.Buffer{}
	words := NewWords(key)

	for {
		w, ok := words.Iter()
		if !ok {
			break
		}
		if len(w) == 0 {
			buf.Reset()
			words.Reset()
			continue
		}

		buf.WriteString(strings.ToLower(w))
		prefix := buf.String()
		if t.t.HasKeysWithPrefix(prefix) {
			setMeta(t.t, res, prefix)
			continue
		}

		buf.Reset()
		words.Reset()
	}

	return res
}

func setMeta(t *gotrie.Trie, res map[string]*Occurring, matchedStr string) {
	if len(matchedStr) == 0 {
		return
	}
	node, ok := t.Find(matchedStr)
	if ok {
		occ, ok := res[matchedStr]
		if !ok {
			res[matchedStr] = &Occurring{
				Count: 1,
				Meta: node.Meta(),
			}
		} else {
			occ.Count += 1
		}
	}
}

