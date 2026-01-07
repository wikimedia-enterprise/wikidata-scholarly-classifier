package scholarly

// Scholarly identifiers (values of P31 = "instance of").
// Based on WDQS scholarly graph split rules.
var scholarlyTypes = map[string]bool{
	"Q13442814": true, // scholarly article
	"Q7318358":  true, // review article
	"Q2782326":  true, // case report
	"Q815382":   true, // meta-analysis
	"Q1348305":  true, // erratum
	"Q187685":   true, // doctoral thesis
	"Q1907875":  true, // master's thesis
	"Q18918145": true, // academic journal article
	"Q1266946":  true, // thesis
	"Q23927052": true, // conference paper
	"Q1504425":  true, // systematic review
	"Q45182324": true, // retracted paper
	"Q580922":   true, // preprint
	"Q15706459": true, // research article
}
