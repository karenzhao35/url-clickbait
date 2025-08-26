package shortener

import (
	cryptorand "crypto/rand"
	"math/big"
	mathrand "math/rand"
)

func genereateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		num, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(charset))))
		shortKey[i] = charset[num.Int64()]
	}
	return string(shortKey)
}

func GenerateCustomKey() string {
	var prefix = []string{
		"hyper", "meta", "quantum", "synergize", "frictionless",
		"scalable", "disruptive", "cloud-native", "lean",
		"fullstack", "nextgen", "tokenized", "ultra", "zero-trust",
	}

	var core = []string{
		"platform", "ecosystem", "pipeline", "protocol", "marketplace",
		"infrastructure", "solution", "metaverse", "dao", "superapp",
		"blockchain", "api-gateway", "crypto-wallet", "datasphere",
		"hoverboard", "dogcoin", "toaster", "bubbletea", "scooter",
	}

	var suffix = []string{
		"-ai", "-as-a-service", "-for-startups", "-dao",
		"-labs", "-ventures", "-capital", "-cloud",
		"-app", "-pro", "-plus", "-xyz", ".io", ".biz",
	}

	randomPrefix := prefix[mathrand.Intn(len(prefix))]
	randomCore := core[mathrand.Intn(len(core))]
	randomSuffix := suffix[mathrand.Intn(len(suffix))]

	return randomPrefix + randomCore + randomSuffix + "-" + genereateShortKey()
}
