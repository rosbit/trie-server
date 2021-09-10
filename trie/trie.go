package trie

var t *Trie

func init() {
	Clear()
}

func Append(key string, meta interface{}) {
	t.Append(key, meta)
}

func Remove(key string) {
	t.Remove(key)
}

func Get(key string) (interface{}, bool) {
	return t.Get(key)
}

func PrefixSearch(key string) []string {
	return t.PrefixSearch(key)
}

func Clear() {
	t = NewTrie()
}

func SimpleSearch(key string) map[string]*Occurring {
	return t.SimpleSearch(key)
}

func Search(key string) map[string]*Occurring {
	return t.Search(key)
}

