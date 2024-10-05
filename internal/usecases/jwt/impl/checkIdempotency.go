package impl

import (
	"context"
)

var (
	blacklistedIdem = make(map[string]bool)
	keyOrder        = []string{}
	maxKeys         = 100
)

func blockIdem(key string) (err error) {
	// Add key to blacklist
	blacklistedIdem[key] = true
	keyOrder = append(keyOrder, key)

	// If the number of keys exceeds maxKeys, remove the oldest key
	if len(keyOrder) > maxKeys {
		oldestKey := keyOrder[0]
		delete(blacklistedIdem, oldestKey)
		keyOrder = keyOrder[1:]
	}

	return
}

func (u *usecase) IsIdemBlacklisted(ctx context.Context, key string) bool {
	_, blacklisted := blacklistedIdem[key]

	err := blockIdem(key)
	if err != nil {
		return true
	}
	return blacklisted
}
