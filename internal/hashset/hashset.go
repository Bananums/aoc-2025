package hashset

import (
	"math"
)

const (
	emptyKey   = math.MaxUint64
	deletedKey = math.MaxUint64 - 1
)

type HashSet struct {
	buckets []uint64
	size    int
}

func NewHashSet(capacity int) *HashSet {
	if capacity < 8 {
		capacity = 8
	}

	capacity = nextPowerOfTwo(capacity)
	buckets := make([]uint64, capacity)

	for i := range buckets {
		buckets[i] = emptyKey
	}

	return &HashSet{
		buckets: buckets,
		size:    0,
	}
}

func nextPowerOfTwo(value int) int {
	value--
	value |= value >> 1
	value |= value >> 2
	value |= value >> 4
	value |= value >> 8
	value |= value >> 16
	if value < 0 {
		return 1
	}
	return value + 1
}

func HashUint64Function(value uint64) uint64 {
	value ^= value >> 30
	value *= 0xbf58476d1ce4e5b9
	value ^= value >> 27
	value *= 0x94d049bb133111eb
	value ^= value >> 31
	return value
}

// Used to check how full the set is
func (hashSet *HashSet) loadFactor() float64 {
	return float64(hashSet.size) / float64(len(hashSet.buckets))
}

// Insert adds key to the set. Returns false if it already was present
func (hashSet *HashSet) Insert(key uint64) bool {

	if key == emptyKey || key == deletedKey {
		panic("hashset: Unsupported key value")
	}

	if hashSet.loadFactor() > 0.75 {
		hashSet.resize(len(hashSet.buckets) * 2)
	}

	mask := uint64(len(hashSet.buckets) - 1)
	hash := HashUint64Function(key)
	index := hash & mask

	var firstDeleted = -1

	for {
		bk := hashSet.buckets[index]
		if bk == emptyKey {
			// Use a previously found deleted slot if any.
			if firstDeleted >= 0 {
				hashSet.buckets[firstDeleted] = key
			} else {
				hashSet.buckets[index] = key
			}
			hashSet.size++
			return true
		}
		if bk == deletedKey {
			// Remember the first tombstone to reuse.
			if firstDeleted < 0 {
				firstDeleted = int(index)
			}
		} else if bk == key {
			// Already present.
			return false
		}
		index = (index + 1) & mask
	}

	return true
}

func (hashSet *HashSet) resize(newCap int) {
	newCap = nextPowerOfTwo(newCap)
	newBuckets := make([]uint64, newCap)
	for i := range newBuckets {
		newBuckets[i] = emptyKey
	}

	oldBuckets := hashSet.buckets
	hashSet.buckets = newBuckets
	hashSet.size = 0

	for _, k := range oldBuckets {
		if k != emptyKey && k != deletedKey {
			hashSet.Insert(k)
		}
	}
}
