package structs

import (
	"math/rand"
	"time"
)

const LIMIT = 25

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

func SpawnSkipList() *SkipList {
	res := &SkipList{
		head: &SkipNode{
			content: 0.0,
			member:  "",
			levels:  make([]*Level, LIMIT),
		},
		level:  1,
		length: 0,
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for i := 0; i < LIMIT; i++ {
		res.head.levels[i] = &Level{
			followingSkip:     nil,
			distanceFollowing: 0,
		}
	}

	res.tail = nil
	return res
}

func CreateSkipNode(level uint, content float64, member string) *SkipNode {
	entity := &SkipNode{
		content: content,
		member:  member,
		levels:  make([]*Level, level),
	}
	for j := 0; j < int(level); j++ {
		entity.levels[j] = &Level{
			followingSkip:     nil,
			distanceFollowing: 0,
		}
	}
	return entity
}
