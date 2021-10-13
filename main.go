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

	trie.insert("dog")

	fmt.Println("Is Ross in the Trie? ", trie.find("ross"))
	fmt.Println("Is Dog in the Trie? ", trie.find("dog"))
	fmt.Println("Is Do in the Trie? ", trie.find("do"))
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
