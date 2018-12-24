package model

type Product struct {
	Name        string
	Description string
	Tags        interface{}
	Category    [][]uint
	MerchanIDs  []uint
	StoresIDs   []uint
	TNC         string
}

type struct Location {
	EntityID uint
	Lat      float64
	Lon      float64
}

type Stores struct {
	ID uint 
	Name uint 
	Image uint
}

	
