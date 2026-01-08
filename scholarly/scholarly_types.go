package scholarly

// Scholarly identifiers (values of P31 = "instance of").
// Based on WDQS scholarly graph split rules.
var scholarlyTypes = map[string]bool{
	"Q13442814":  true, // scholarly article
	"Q7318358":   true, // review article
	"Q2782326":   true, // case report
	"Q815382":    true, // meta-analysis
	"Q1348305":   true, // erratum
	"Q187685":    true, // doctoral thesis
	"Q1907875":   true, // master's thesis
	"Q18918145":  true, // academic journal article
	"Q1266946":   true, // thesis
	"Q23927052":  true, // conference paper
	"Q1504425":   true, // systematic review
	"Q45182324":  true, // retracted paper
	"Q1402850":   true, // field study report
	"Q7316896":   true, // retraction notice
	"Q580922":    true, // preprint
	"Q30749496":  true, // diploma thesis
	"Q111475835": true, // bachelor's with honors thesis
	"Q92998777":  true, // opinion paper
	"Q114613919": true, // scientific note
	"Q798134":    true, // bachelor’s thesis
	"Q1385450":   true, // dissertation
	"Q10885494":  true, // academic conference paper
	"Q51282918":  true, // Doctor of Philosophy thesis
	"Q51282711":  true, // Doctor of Clinical Psychology thesis
	"Q111475860": true, // postgraduate diploma thesis
	"Q51283092":  true, // Master of Arts thesis
	"Q15706459":  true, // research article
	"Q59387148":  true, // research report
	"Q110716513": true, // scholarly letter/reply
	"Q58897583":  true, // comment
	"Q51283145":  true, // Doctor of Medicine thesis
	"Q54670950":  true, // conference poster
	"Q91901000":  true, // multicenter study report
	"Q111476177": true, // masters with honours thesis
	"Q51283053":  true, // Doctor of Education thesis
	"Q1414362":   true, // habilitation thesis
	"Q51283181":  true, // Master of Philosophy thesis
	"Q51282999":  true, // Doctor of Science thesis
	"Q51283199":  true, // Master of Science thesis
	"Q82969330":  true, // medical scholarly article
	"Q112585758": true, // Bachelor of Literature
	"Q118114827": true, // postgraduate certificate thesis
	"Q106276531": true, // bachelor’s thesis
	"Q1884156":   true, // Magister's thesis
	"Q51283362":  true, // Master of Letters or Literature thesis
	"Q46629343":  true, // Archivist palaeographer thesis
	"Q100328456": true, // doctoral thesis (sciences)
	"Q51283219":  true, // Master of Science by Research thesis
	"Q70471362":  true, // non-randomized controlled trial report
}
