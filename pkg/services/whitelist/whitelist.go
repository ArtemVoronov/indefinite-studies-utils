package whitelist

import (
	"os"
	"regexp"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
)

type WhiteListService struct {
	store map[string]string
}

func CreateWhiteListService(path string) *WhiteListService {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to create white list service: %v", err)
	}
	input := string(data[:])

	re := regexp.MustCompile("\n")
	tokens := re.Split(input, -1)
	store := make(map[string]string)
	for _, line := range tokens {
		store[line] = ""
	}
	return &WhiteListService{store}
}

func (s *WhiteListService) Shutdown() error {
	return nil
}

func (s *WhiteListService) Contains(key string) bool {
	_, ok := s.store[key]
	return ok
}
