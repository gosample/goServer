package models

// Result as a return json template.
type RegisterResult struct {
	Result int
	Err    string
}

type LoginResult struct {
	Result int
	Err    string
	Token  string
}
type SearchResult struct {
	Result int
	Err    string
	Parks  []Result
}
type NearByResult struct {
	Result int
	Err    string
	Parks  []Result
}
