package dto

//LatLong lat and long struct
type LatLong struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}

//Device Information struct
type Device struct {
	ID           string `json:"id"`
	OS           string `json:"os"`
	OsVersion    string `json:"osVersion"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
}
