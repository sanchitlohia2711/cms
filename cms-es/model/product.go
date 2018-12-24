package model

//ProductES struct
type ProductES struct {
	Name        string
	Description string
	Tags        interface{}
	Category    [][]uint
	MerchanIDs  []uint
	StoresIDs   []uint
	TNC         string
}

type Location struct {
	EntityID uint
	Lat      float64
	Lon      float64
}

type Stores struct {
	ID    uint
	Name  uint
	Image uint
}
