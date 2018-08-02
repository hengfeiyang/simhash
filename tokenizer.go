package simhash

// Tokenizer word tokenizer
type Tokenizer struct {
	chunkSize, overlapSize uint8
}

// TokenizerChunk word chunk
type TokenizerChunk struct {
	Word   string
	Weight int
}

// NewTokenizer create a new tokenizer
// chunkSize,   suggestion value: 4
// overlapSize, suggestion value: 1
func NewTokenizer(chunkSize, overlapSize uint8) *Tokenizer {
	if chunkSize <= overlapSize {
		panic("chunk size must be greater than overlap size.")
	}

	return &Tokenizer{chunkSize: chunkSize, overlapSize: overlapSize}
}

// Tokenize execute tokenize
// simple set words weight value: 1
func (t *Tokenizer) Tokenize(input string) []TokenizerChunk {
	var chunks []TokenizerChunk
	inputRune := []rune(input)
	inputLen := len(inputRune)
	for position := 0; position < inputLen-int(t.chunkSize); position += int(t.chunkSize - t.overlapSize) {
		chunks = append(chunks, TokenizerChunk{string(inputRune[position : position+int(t.chunkSize)]), 1})
	}
	return chunks
}
