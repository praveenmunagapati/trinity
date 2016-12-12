package subsystem

import "github.com/clownpriest/trinity/api/trinity"

/*
Search sends a search query to the backend systems that will
process it and return a search result
*/
func Search(query *trinity.SearchQuery) *trinity.SearchResponse {
	results := SampleSearchResults()
	response := &trinity.SearchResponse{Results: results}
	return response
}

/*
SampleSearchResults just returns some dummy results for testing
purposes
*/
func SampleSearchResults() []*trinity.SearchResult {
	return []*trinity.SearchResult{

		&trinity.SearchResult{
			Title:       "Gnosis - Wikipedia",
			Hash:        "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			Description: "Gnosis is the common Greek noun for knowledge (in the nominative case γνῶσις f.). It generally signifies a dualistic knowledge in the sense of mystical enlightenment or \"insight\".",
		},
		&trinity.SearchResult{
			Title:       "Artificial Intelligence - Wikipedia",
			Hash:        "QmPz0f3zv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			Description: "Artificial intelligence (AI) is intelligence exhibited by machines. In computer science, an ideal \"intelligent\" machine is a flexible rational agent that perceives its ...",
		},
		&trinity.SearchResult{
			Title:       "Isomorphism - Wikipedia",
			Hash:        "QmTo3Guzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			Description: "In mathematics, an isomorphism (from the Ancient Greek: ἴσος isos \"equal\", and μορφή morphe \"form\" or \"shape\") is a homomorphism or morphism (i.e. a mathematical mapping) that admits an inverse.",
		},
		&trinity.SearchResult{
			Title:       "Computer Science - Wikipedia",
			Hash:        "QmW4j0pzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			Description: "An alternate, more succinct definition of computer science is the study of automating algorithmic processes that scale. A computer scientist specializes in the theory of computation and the design of computational systems. Its fields can be divided into a variety of theoretical and practical disciplines.",
		},
		&trinity.SearchResult{
			Title:       "Google - Wikipedia",
			Hash:        "QmJl3eYbv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG",
			Description: "Google is an American multinational technology company specializing in Internet-related services and products that include online advertising technologies, ...",
		},
	}

}
