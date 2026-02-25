package service

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/brmcoder/line-bot-dict/model"
	"github.com/iancoleman/strcase"
)

type dictionaryService interface {
	GetWordDefinition(word string) (string, error)
}

type dictionaryServ struct{}

// GetWordDefinition implements dictionaryService.
func (dictionaryServ) GetWordDefinition(word string) (string, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound {
		return "", fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	if resp.StatusCode == http.StatusNotFound {
		return "No definitions found.", nil
	}

	var result []model.DefinitionResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	var response string
	for i, meaning := range result[0].Meanings {
		if i == len(result[0].Meanings)-1 {
			response += fmt.Sprintf("■ %s: %s", strcase.ToCamel(meaning.PartOfSpeech), meaning.Definitions[0].Definition)
		} else {
			response += fmt.Sprintf("■ %s: %s\n", strcase.ToCamel(meaning.PartOfSpeech), meaning.Definitions[0].Definition)
		}
	}

	return response, nil
}

func NewDictionaryService() dictionaryService {
	return dictionaryServ{}
}
