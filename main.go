package main

import "fmt"

const (
	AlphabetSize = 26
)

type trie struct {
	root *node
}

type node struct {
	children  [AlphabetSize]*node
	endOfWord bool
}

func main() {
	trie := createTrie()

	// TODO When this runs we need to popualte this from a Source
	trie.insert("dos")
	trie.insert("doggy")
	trie.insert("domination")

	// TODO Turn this into a HTTP server to get suggestions based off a query string
	dogSuggestions := trie.suggest("do")
	fmt.Println("Printing Suggestions for 'do'")
	for _, s := range dogSuggestions {
		fmt.Println(s)
	}

}

func (t *trie) suggest(word string) []string {
	var suggestions []string
	wordLength := len(word)

	currentNode := t.root
	var prefix string

	for i := 0; i < wordLength; i++ {
		letterIndex := word[i] - 'a'
		prefix = fmt.Sprintf("%s%s", prefix, string(word[i]))

		// We have a Word that is not in the Trie, return an empty suggestions slice.
		if currentNode.children[letterIndex] == nil {
			return suggestions
		}

		currentNode = currentNode.children[letterIndex]
	}

	// Check if this is an end node (nothing to suggest)
	// We may want to think about if it is end of word too, for example "dog" is end of word
	// but do we want to suggest "doggy"?
	if currentNode.isLastLeaf() {
		return suggestions
	}

	// We need to Recursively get the children of this node, and add them to suggestions.
	// We probably need to think about Limits, Getting Popular items (total searches?!) and goroutines.
	// I can see this getting out of hand for smaller queries with a lot of words
	for i := 0; i < AlphabetSize; i++ {
		tmpWord := fmt.Sprintf("%s%s", prefix, string(i+'a'))
		if currentNode.children[i] != nil {
			w := currentNode.children[i].suggestWord(tmpWord)
			if w != "" {
				suggestions = append(suggestions, w)
			}
		}
	}

	return suggestions
}

// suggestWord will return a string after recursively going through the children of the called on node
func (n *node) suggestWord(prefix string) string {
	if n.endOfWord || n.isLastLeaf() {
		return prefix
	}

	// For each child node we need to recursively call this method
	for i := 0; i < AlphabetSize; i++ {
		if n.children[i] != nil {
			prefix = fmt.Sprintf("%s%s", prefix, string(i+'a'))
			w := n.children[i].suggestWord(prefix)
			if w != "" {
				return w
			}
		}
	}

	// In theory this should never reach here because when you go into a node, it should either be the
	// end of a word or the last node (also should be end of word), so we should never not return
	// something from the for loop
	return ""
}

// insert takes a string and adds it to the root trie
func (t *trie) insert(word string) {
	wordLength := len(word)

	currentNode := t.root

	// Foreach Letter in the Word we need to check the next Trie Node
	for i := 0; i < wordLength; i++ {
		letterIndex := word[i] - 'a'

		// We need to add a new node in at this letter index
		if currentNode.children[letterIndex] == nil {
			currentNode.children[letterIndex] = createNode()
		}

		// Set the current node to the child node that we have or that has been created
		currentNode = currentNode.children[letterIndex]
	}

	// Set end of word
	currentNode.endOfWord = true
}

// find takes the root trie and a word and returns a bool as to whether the
// string is found within the trie or not
func (t *trie) find(word string) bool {
	wordLength := len(word)

	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		letterIndex := word[i] - 'a'

		// We don't have an index here so the word isn't in the trie
		if currentNode.children[letterIndex] == nil {
			return false
		}

		currentNode = currentNode.children[letterIndex]
	}

	return currentNode.endOfWord
}

// isLastLeaf will loop through all the children nodes and return based on if there is atleast one node
func (n *node) isLastLeaf() bool {
	for i := 0; i < AlphabetSize; i++ {
		if n.children[i] != nil {
			return false
		}
	}

	return true
}

// createTrie creates the root trie with an initial node
func createTrie() *trie {
	return &trie{
		root: createNode(),
	}
}

// createNode creates and returns a basic node
func createNode() *node {
	return &node{
		endOfWord: false,
	}
}
