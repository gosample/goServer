package models

// Result as a return json template.
type RegisterResult struct {
	Result int64
	Err    string
}

type LoginResult struct {
	Result      int64
	Err         string
	Token       string
	CarLicenses []CarResult
}
type SearchResult struct {
	Result int64
	Err    string
	Parks  []ParkResult
}
type NearByResult struct {
	Result int64
	Err    string
	Parks  []ParkResult
}
type BookResult struct {
	Result int64
	Err    string
}
type AddCarLicenseResult struct {
	Result int64
	Err    string
}
type GetParkInfoResult struct {
	Result    int64
	Err       string
	StoreyNum int64
	EmptyNum  int64
	Money     float64
	ParkImg   string
}
type ServiceResult struct {
	Result   int64
	Err      string
	Services []Service
}

type SpaceResult struct {
	Result     int64
	Num        int64
	Err        string
	ParkSpaces []ParkSpace
}
