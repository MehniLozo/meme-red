package ops

import (
	types "github.com/MehniLozo/meme-red/structs/types"
)

func Int[T types.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
