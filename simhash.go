package simhash

import (
	"hash/fnv"
	"io"
)

const (
	// HashSize simhash length，choose 32/64
	HashSize = 32
)

// Simhash calc given string simhash
func Simhash(input string) uint64 {
	tokeniser := NewTokenizer(4, 3)
	hashedTokens := getHashTokens(tokeniser.Tokenize(input))
	vector := make([]int, HashSize)

	for _, v := range hashedTokens {
		for i := range vector {
			if isBitSet(v.Hash, uint64(i)) {
				vector[i] += v.Weight
			} else {
				vector[i] -= v.Weight
			}
		}
	}

	var fingerprint uint64
	for i, v := range vector {
		if v > 0 {
			fingerprint += 1 << uint8(i)
		}
	}

	return fingerprint
}

// Distance calc two simhash haiming distance
func Distance(a, b uint64) int {
	hammingBits := a ^ b
	hammingValue := 0
	for i := uint64(0); i < HashSize; i++ {
		if isBitSet(hammingBits, i) {
			hammingValue++
		}
	}
	return hammingValue
}

// Similar calc two simhash similar percent
func Similar(a, b uint64) float64 {
	return float64(HashSize-Distance(a, b)) / float64(HashSize)
}

type hashToken struct {
	Hash   uint64
	Weight int
}

// getHashTokens calc tokens hash
// can choose hash algorithm, hash/crc32, hash/crc64, has/fnv...
func getHashTokens(tokens []TokenizerChunk) []hashToken {
	hashedTokens := make([]hashToken, len(tokens))
	f := fnv.New32()
	for i, token := range tokens {
		f.Reset()
		io.WriteString(f, token.Word)
		hashedTokens[i] = hashToken{uint64(f.Sum32()), token.Weight}
	}
	return hashedTokens
}

func isBitSet(b, pos uint64) bool {
	return (b & (1 << pos)) != 0
}
