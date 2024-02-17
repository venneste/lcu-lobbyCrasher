package lcu

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func makeLCUHeader(authToken string) http.Header {
	sessionToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("riot:%s", authToken)))

	return http.Header{
		"Content-Type":  {"application/json"},
		"Accept":        {"application/json"},
		"Authorization": {"Basic " + sessionToken + ""},
	}
}

func TrimQuotes(str string) string {
	return strings.Trim(str, "\"")
}

func NewLeagueSessionToken(responseString string) *LeagueSessionToken {
	jwt := LeagueSessionToken(TrimQuotes(responseString))

	return &jwt
}

func (r *LeagueSessionToken) String() string {
	return string(*r)
}

func NewRSOInventoryJWT(responseString string) *RSOInventoryJWT {
	jwt := RSOInventoryJWT(TrimQuotes(responseString))

	return &jwt
}

func (r *RSOInventoryJWT) String() string {
	return string(*r)
}

func MakeGameMap() *GameMap {
	return &GameMap{
		Class:            "com.riotgames.platform.game.map.GameMap",
		Description:      "",
		DisplayName:      "",
		MapId:            11,
		MinCustomPlayers: 1,
		Name:             "",
		TotalPlayers:     10,
	}
}
