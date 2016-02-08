package model

type ConceptSuggestion struct {
	UUID        string       `json:"uuid"`
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Thing      Thing        `json:"thing"`
	Provenance []Provenance `json:"provenances"`
}

type Thing struct {
	Id        string   `json:"id"`
	PrefLabel string   `json:"prefLabel"`
	Types     []string `json:"types"`
}

type Provenance struct {
	Scores []Score `json:"scores"`
}

type Score struct {
	ScoringSystem string  `json:"scoringSystem"`
	Value         float32 `json:"value"`
}
