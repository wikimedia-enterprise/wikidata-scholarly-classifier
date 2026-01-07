package scholarly

import "strings"

// IsScholarly determines whether the given Wikidata entity
// belongs to the scholarly graph.
func IsScholarly(entity *WikidataEntity) bool {
	if entity == nil || entity.Statements == nil {
		return false
	}

	// Rule 1 — publication type of scholarly work (P13046)
	if stmts, ok := entity.Statements["P13046"]; ok {
		for _, s := range stmts {
			if !isDeprecated(s.Rank) {
				return true
			}
		}
	}

	// Rule 2 — instance of (P31)
	if stmts, ok := entity.Statements["P31"]; ok {
		for _, s := range stmts {
			if isDeprecated(s.Rank) {
				continue
			}
			if scholarlyTypes[extractStatementValueID(&s)] {
				return true
			}
		}
	}

	return false
}

func isDeprecated(rank string) bool {
	return strings.EqualFold(rank, "deprecated")
}

func extractStatementValueID(stmt *WikidataStatement) string {
	if stmt == nil {
		return ""
	}

	val := stmt.Value
	if val.Type == "wikibase-entityid" {
		if id, ok := val.ContentData["id"]; ok && id != "" {
			return id
		}
		return val.Content
	}

	return ""
}
