package chainedmap

type ChainedMap[K comparable, V any] struct {
	maps []map[K]V
}

func New[K comparable, V any](base map[K]V) *ChainedMap[K, V] {
	m := new(ChainedMap[K, V])
	if base == nil {
		base = make(map[K]V)
	}

	m.maps = append(m.maps, base)

	return m
}

// Get searches for the value in the down-top direction
func (c *ChainedMap[K, V]) Get(key K) (value V, found bool) {
	for i := len(c.maps) - 1; i >= 0; i-- {
		if value, found = c.maps[i][key]; found {
			return value, true
		}
	}

	return value, false
}

// Insert inserts a new value to the down. Old value may be overridden
func (c *ChainedMap[K, V]) Insert(key K, value V) {
	c.maps[len(c.maps)-1][key] = value
}

// Update looks for the value in the down-top direction, and overrides its value only
// in case it exists somewhere
func (c *ChainedMap[K, V]) Update(key K, value V) (updated bool) {
	for i := len(c.maps) - 1; i >= 0; i-- {
		if _, found := c.maps[i][key]; found {
			c.maps[i][key] = value
			return true
		}
	}

	return false
}

// Pop removes one level of nesting
func (c *ChainedMap[K, V]) Pop() {
	if len(c.maps) == 0 {
		return
	}

	c.maps = c.maps[:len(c.maps)-1]
}

// Push creates a new level of nesting
func (c *ChainedMap[K, V]) Push() {
	c.maps = append(c.maps, make(map[K]V))
}
