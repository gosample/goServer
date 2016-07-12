package models

// Result as a return json template.
type RegisterResult struct {
	Result int
	Err    string
}

type LoginResult struct {
	Result      int
	Err         string
	Token       string
	CarLicenses []CarResult
}
type SearchResult struct {
	Result int
	Err    string
	Parks  []ParkResult
}
type NearByResult struct {
	Result int
	Err    string
	Parks  []ParkResult
}
type BookResult struct {
	Result int
	Err    string
}
type AddCarLicenseResult struct {
	Result int
	Err    string
}
type GetParkInfoResult struct {
	Result    int
	Err       string
	StoreyNum int
	EmptyNum  int
	Money     float64
	ParkImg   string
}
type ServiceResult struct {
	Result   int
	Err      string
	Services []Service
}
