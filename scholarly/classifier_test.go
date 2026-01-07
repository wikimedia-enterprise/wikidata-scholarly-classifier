package scholarly

import "testing"

func TestIsScholarly_P31ScholarlyType(t *testing.T) {
	entity := &WikidataEntity{
		Statements: map[string][]WikidataStatement{
			"P31": {{
				Rank: "normal",
				Value: WikidataValue{
					Type: "wikibase-entityid",
					ContentData: map[string]string{
						"id": "Q13442814",
					},
				},
			}},
		},
	}

	if !IsScholarly(entity) {
		t.Fatalf("expected entity to be scholarly")
	}
}

func TestIsScholarly_P13046(t *testing.T) {
	entity := &WikidataEntity{
		Statements: map[string][]WikidataStatement{
			"P13046": {{
				Rank: "normal",
			}},
		},
	}

	if !IsScholarly(entity) {
		t.Fatalf("expected entity to be scholarly due to P13046")
	}
}

func TestIsScholarly_DeprecatedStatementsIgnored(t *testing.T) {
	entity := &WikidataEntity{
		Statements: map[string][]WikidataStatement{
			"P31": {{
				Rank: "deprecated",
				Value: WikidataValue{
					Type: "wikibase-entityid",
					ContentData: map[string]string{
						"id": "Q13442814",
					},
				},
			}},
		},
	}

	if IsScholarly(entity) {
		t.Fatalf("expected entity to be non-scholarly due to deprecated statement")
	}
}

func TestIsScholarly_NilEntity(t *testing.T) {
	if IsScholarly(nil) {
		t.Fatalf("expected nil entity to be non-scholarly")
	}
}
