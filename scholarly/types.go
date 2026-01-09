package scholarly

// WikidataEntity is a minimal representation of a Wikidata entity,
// containing only the data required for scholarly classification.
type WikidataEntity struct {
	Statements map[string][]WikidataStatement
}

// WikidataStatement is a minimal statement representation.
type WikidataStatement struct {
	Rank  string
	Value WikidataValue
}

// WikidataValue is a minimal value representation.
type WikidataValue struct {
	Type        string
	Content     string
	ContentData map[string]string
}
