package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"os"
	"strconv"
	"time"

	"github.com/ArtemVoronov/indefinite-studies-utils/pkg/log"
)

func CreateSHA512HashHexEncoded(str string) string {
	hasher := sha512.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateSHA256HashHexEncoded(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Contains(array []string, elem string) bool {
	for _, n := range array {
		if elem == n {
			return true
		}
	}
	return false
}

func EnvVar(varName string) string {
	val, valExists := os.LookupEnv(varName)
	if !valExists {
		log.Fatalf("Missed enviroment variable: %s. Check the .env file or OS enviroment vars", varName)
	}
	return val
}

func EnvVarDefault(varName string, defaultValue string) string {
	val, valExists := os.LookupEnv(varName)
	if !valExists {
		return defaultValue
	}
	return val
}

func EnvVarInt(varName string) int {
	val := EnvVar(varName)
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Wrong value of environment variable: %s. It should be integer number", varName)
	}
	return result
}

func EnvVarIntDefault(varName string, defaultValue int) int {
	val, valExists := os.LookupEnv(varName)
	if !valExists {
		return defaultValue
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Wrong value of environment variable: %s. It should be integer number", varName)
	}
	return result
}

func EnvVarDuration(varName string, unit time.Duration) time.Duration {
	val := EnvVarInt(varName)
	return unit * time.Duration(val)
}

func EnvVarDurationDefault(varName string, unit time.Duration, defaultValue time.Duration) time.Duration {
	val, valExists := os.LookupEnv(varName)
	if !valExists {
		return defaultValue
	}
	number, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Wrong value of environment variable: %s. It should be integer number", varName)
	}
	return unit * time.Duration(number)
}

func EnvVarBytes(varName string) []byte {
	val := EnvVar(varName)
	return []byte(val)
}

func Int32ToIntPtr(val int32) *int {
	if val == 0 {
		return nil
	}
	result := int(val)
	return &result
}

func IntPtrToInt32(val *int) int32 {
	if val == nil {
		return 0
	}
	result := int32(*val)
	return result
}

func Int32SliceToIntSlice(input []int32) []int {
	result := make([]int, len(input))
	for i := range result {
		result[i] = int(input[i])
	}
	return result
}

func ToInt32(in []int) []int32 {
	result := make([]int32, 0, len(in))

	for _, p := range in {
		result = append(result, int32(p))
	}

	return result
}

func ToInt(in []int64) []int {
	result := make([]int, 0, len(in))

	for _, p := range in {
		result = append(result, int(p))
	}

	return result
}
