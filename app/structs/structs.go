package structs

import (
	"math/rand"
)

const MAX = 25

type SkipNode struct {
	member  string
	content float64
	levels  []*Level
}

type SkipList struct {
	head, tail *SkipNode
	length     uint64
	level      uint
	random     rand.Source
}

type Level struct {
	followingSkip     *SkipNode
	distanceFollowing uint64
}
