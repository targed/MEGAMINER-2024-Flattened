package magomachy

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Object is the Item on this Tile if present, otherwise nil.
	//
	// Value can be returned as a nil pointer.
	Object() Item

	// TileEast is the Tile to the 'East' of this one (x+1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileEast() Tile

	// TileNorth is the Tile to the 'North' of this one (x, y-1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileNorth() Tile

	// TileSouth is the Tile to the 'South' of this one (x, y+1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileSouth() Tile

	// TileWest is the Tile to the 'West' of this one (x-1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileWest() Tile

	// Type is the type of Tile this is (i.e Grass, Wall).
	//
	// Literal Values: "floor" or "wall"
	Type() string

	// Wizard is the Wizard on this Tile if present, otherwise nil.
	//
	// Value can be returned as a nil pointer.
	Wizard() Wizard

	// X is the x (horizontal) position of this Tile.
	X() int64

	// Y is the y (vertical) position of this Tile.
	Y() int64

	// -- Tiled Game Utils -- \\

	// GetNeighbors returns an array of the neighbors of this Tile.
	GetNeighbors() []Tile

	// IsPathable returns if the Tile is pathable for FindPath
	IsPathable() bool

	// HasNeighbor checks if this Tile has a specific neighboring Tile.
	HasNeighbor(Tile) bool
}

// TileDirections are all the direction strings that Tile's neighbors can be
// in.
var TileDirections = [...]string{"North", "South", "East", "West"}
