package trie

// Trie represents a trie holding all prefixes
type Node struct {
	// IsPrefEnd flags if the node holds last char of a prefix
	IsPrefEnd bool
	// map of matching prefiex in this node
	Children map[string]*Node
}

// NewNode returns the pointer to a new node
// sets Children map and IsPrefEnd flag to thier zero values
func NewNode() *Node {
	var tn Node

	// initialize the map
	tn.Children = make(map[string]*Node)
	tn.IsPrefEnd = false

	// return the pointer to the trie node
	return &tn
}

// Build builds the trie with each char in the string
func (tr *Node) Build(prefix string) {
	if len(prefix) == 0 {
		return
	}

	cur := tr

	for _, c := range prefix {
		c := string(c)
		nn, ok := cur.Children[c]
		if !ok {
			nn = NewNode()
			cur.Children[c] = nn
		}
		cur = nn
	}

	cur.IsPrefEnd = true
}

// SearchLongestPrefix returns the longest prefix matched for the given search word
// returns empty string if not found
func (tr *Node) SearchLongestPrefix(searchWord string) string {

	var lastPref string
	var iMatchStr string

	cur := tr

	for _, ch := range searchWord {
		ch := string(ch)
		if cur.Children[ch] != nil {
			iMatchStr += ch
			if cur.Children[ch].IsPrefEnd == true {
				lastPref += iMatchStr
				iMatchStr = ""
			}
			cur = cur.Children[ch]
		}
	}

	return lastPref
}
