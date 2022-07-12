package game

type Tile rune

const (
	StoneWall Tile = '#'
	DirtFloor Tile = '.'
	Door      Tile = '|'
)

type level struct {
	Map [][]Tile
}
