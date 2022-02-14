package form3

type Links struct {
	// Link to the first resource in the list
	First string `json:"first,omitempty"`
	// Link to the last resource in the list
	Last string `json:"last,omitempty"`
	// Link to the next resource in the list
	Next string `json:"next,omitempty"`
	// Link to the previous resource in the list
	Prev string `json:"prev,omitempty"`
	// Link to this resource type
	Self string `json:"self"`
}
