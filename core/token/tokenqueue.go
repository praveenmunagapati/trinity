package token

/*
Queue is a list of token.List
*/
type Queue struct {
	data []List
}

/*
NewQueue returns a new empty Queue
*/
func NewQueue() *Queue {
	return &Queue{}
}

// Add appends a token.List to the queue
func (q *Queue) Add(list List) {
	q.data = append(q.data, list)
}

/*
Clear releases the data, making it eligible for
garbage collection
*/
func (q *Queue) Clear() {
	q.data = []List{}
}

/*
Data returns the underlying data in the Queue
*/
func (q *Queue) Data() []List {
	return q.data
}
