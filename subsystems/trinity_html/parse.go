package thtml

import (
	"os"
	"path"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/clownpriest/trinity/core/intake"
	"github.com/clownpriest/trinity/core/token"
)

func traverse(node *html.Node, buffer *intake.BufferPool) {
	if uselessNodeType(node) {
		return
	}

	if node.Type == html.TextNode {
		buffer.WriteString(node.Data)
	}

	traverseChildren(node, buffer)

}

func traverseChildren(node *html.Node, buffer *intake.BufferPool) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverse(child, buffer)
	}
}

func uselessNodeType(node *html.Node) bool {
	if node.DataAtom == atom.Style || node.DataAtom == atom.Script {
		return true
	}
	return false
}

/*
Tokenize returns the contents of a byte buffer
in string token form.
*/
func Tokenize(buffer *intake.BufferPool, filename string) token.List {
	str := buffer.String()
	tokens := FilterPipeline(str)
	tokens.SetFilename(path.Base(filename))
	return tokens
}

/*
Read parses an HTML file into an intake.Buffer
*/
func Read(path string, buffer *intake.BufferPool) {
	file, _ := os.Open(path)

	doc, err := html.Parse(file)
	if err != nil {
	}

	traverse(doc, buffer)
}

/*
ReadLoop takes a buffer pointer and a list of paths. It reads these
paths into the buffer in a loop, tokenizes their content, and clears
the buffer for the next HTML file in the path list. Each tokenized output
if pushed onto a token.List Queue.
*/
func ReadLoop(paths *intake.FileList, buffer *intake.BufferPool, queue *token.Queue) {
	for _, path := range *paths {
		file, _ := os.Open(path)
		doc, err := html.Parse(file)
		if err != nil {
		}

		traverse(doc, buffer)

		tokens := Tokenize(buffer, path)
		queue.Add(tokens)
		//fmt.Println(queue)
		buffer.Reset()
	}
}
