package strand

var complements = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

//ToRNA - transcribes the RNA for a given a DNA strand
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, nucleotide := range dna {
		rna[i] = complements[nucleotide]
	}
	return string(rna)
}
