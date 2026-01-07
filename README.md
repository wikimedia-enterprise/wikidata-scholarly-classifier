# Wikidata Scholarly Classifier

A lightweight Go library for classifying Wikidata entities as **scholarly** or **non-scholarly** based on Wikidata statements.

The logic follows the WDQS scholarly graph split rules.
## Rules

An entity is considered scholarly if:
- It has a non-deprecated **publication type of scholarly work** (`P13046`), or
- It has a non-deprecated **instance of** (`P31`) whose value is a known scholarly type.

## Usage

```go
import "github.com/wikimedia/wikidata-scholarly-classifier/scholarly"

if scholarly.IsScholarly(entity) {
    // handle scholarly entity
}
