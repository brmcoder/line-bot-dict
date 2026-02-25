package model

type DefinitionResult struct {
	Word       string    `json:"word"`
	Meanings   []Meaning `json:"meanings"`
	SourceUrls []string  `json:"sourceUrls"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
	Example    *string  `json:"example,omitempty"`
}
