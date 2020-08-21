package protein

import (
	"errors"
)

var (
	ErrInvalidBase = errors.New("Invalid base")
	ErrStop        = errors.New("Stop error")
)

//FromCodon - From codon to protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

//FromRNA - From protein to codon
func FromRNA(rna string) ([]string, error) {
	codonLength := 3
	var result []string
	for i := 0; i < len(rna)/3; i++ {
		protein, err := FromCodon(rna[i*codonLength : (i+1)*codonLength])
		if err == ErrStop {
			break
		}
		if err != nil {
			return result, err
		}
		result = append(result, protein)
	}
	return result, nil
}
