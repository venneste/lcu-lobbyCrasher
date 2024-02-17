package lcu

import (
	"github.com/ImOlli/go-lcu/lcu"
	"net/http"
)

type LCU struct {
	DestURL string
	Info    *lcu.ConnectInfo
	Client  *http.Client
}

type RSOIdToken struct {
	Expiry int64  `json:"expiry"`
	Token  string `json:"token"`
}

type RSOAccessToken struct {
	Expiry int      `json:"expiry"`
	Scopes []string `json:"scopes"`
	Token  string   `json:"token"`
}

type LobbySummoner struct {
	AssignedPosition     string `json:"assignedPosition"`
	CellId               int    `json:"cellId"`
	ChampionId           int    `json:"championId"`
	ChampionPickIntent   int    `json:"championPickIntent"`
	NameVisibilityType   string `json:"nameVisibilityType"`
	ObfuscatedPuuid      string `json:"obfuscatedPuuid"`
	ObfuscatedSummonerId int    `json:"obfuscatedSummonerId"`
	Puuid                string `json:"puuid"`
	SelectedSkinId       int    `json:"selectedSkinId"`
	Spell1Id             int    `json:"spell1Id"`
	Spell2Id             int    `json:"spell2Id"`
	SummonerId           int64  `json:"summonerId"`
	Team                 int    `json:"team"`
	WardSkinId           int    `json:"wardSkinId"`
}

type Timer struct {
	AdjustedTimeLeftInPhase int    `json:"adjustedTimeLeftInPhase"`
	InternalNowInEpochMs    int64  `json:"internalNowInEpochMs"`
	IsInfinite              bool   `json:"isInfinite"`
	Phase                   string `json:"phase"`
	TotalTimeInPhase        int    `json:"totalTimeInPhase"`
}

type ChampSelectSession struct {
	Actions [][]struct {
		ActorCellId  int    `json:"actorCellId"`
		ChampionId   int    `json:"championId"`
		Completed    bool   `json:"completed"`
		Id           int    `json:"id"`
		IsAllyAction bool   `json:"isAllyAction"`
		IsInProgress bool   `json:"isInProgress"`
		PickTurn     int    `json:"pickTurn"`
		Type         string `json:"type"`
	} `json:"actions"`
	AllowBattleBoost    bool `json:"allowBattleBoost"`
	AllowDuplicatePicks bool `json:"allowDuplicatePicks"`
	AllowLockedEvents   bool `json:"allowLockedEvents"`
	AllowRerolling      bool `json:"allowRerolling"`
	AllowSkinSelection  bool `json:"allowSkinSelection"`
	Bans                struct {
		MyTeamBans    []interface{} `json:"myTeamBans"`
		NumBans       int           `json:"numBans"`
		TheirTeamBans []interface{} `json:"theirTeamBans"`
	} `json:"bans"`
	BenchChampions     []interface{} `json:"benchChampions"`
	BenchEnabled       bool          `json:"benchEnabled"`
	BoostableSkinCount int           `json:"boostableSkinCount"`
	ChatDetails        struct {
		MucJwtDto struct {
			ChannelClaim string `json:"channelClaim"`
			Domain       string `json:"domain"`
			Jwt          string `json:"jwt"`
			TargetRegion string `json:"targetRegion"`
		} `json:"mucJwtDto"`
		MultiUserChatId       string `json:"multiUserChatId"`
		MultiUserChatPassword string `json:"multiUserChatPassword"`
	} `json:"chatDetails"`
	Counter              int             `json:"counter"`
	GameId               int             `json:"gameId"`
	HasSimultaneousBans  bool            `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool            `json:"hasSimultaneousPicks"`
	IsCustomGame         bool            `json:"isCustomGame"`
	IsSpectating         bool            `json:"isSpectating"`
	LocalPlayerCellId    int             `json:"localPlayerCellId"`
	LockedEventIndex     int             `json:"lockedEventIndex"`
	MyTeam               []LobbySummoner `json:"myTeam"`
	PickOrderSwaps       []interface{}   `json:"pickOrderSwaps"`
	RecoveryCounter      int             `json:"recoveryCounter"`
	RerollsRemaining     int             `json:"rerollsRemaining"`
	SkipChampionSelect   bool            `json:"skipChampionSelect"`
	TheirTeam            []LobbySummoner `json:"theirTeam"`
	Timer                Timer           `json:"timer"`
	Trades               []interface{}   `json:"trades"`
}

type GameMap struct {
	Class            string `json:"__class"`
	Description      string `json:"description"`
	DisplayName      string `json:"displayName"`
	MapId            int    `json:"mapId"`
	MinCustomPlayers int    `json:"minCustomPlayers"`
	Name             string `json:"name"`
	TotalPlayers     int    `json:"totalPlayers"`
}

type PracticeGameConfig struct {
	Class              string      `json:"__class"`
	AllowSpectators    string      `json:"allowSpectators"`
	GameMap            GameMap     `json:"gameMap"`
	GameMode           string      `json:"gameMode"`
	GameMutators       []int       `json:"gameMutators"`
	GameName           string      `json:"gameName"`
	GamePassword       string      `json:"gamePassword"`
	GameTypeConfig     int         `json:"gameTypeConfig"`
	GameVersion        string      `json:"gameVersion"`
	MaxNumPlayers      int         `json:"maxNumPlayers"`
	PassbackDataPacket interface{} `json:"passbackDataPacket"`
	PassbackUrl        interface{} `json:"passbackUrl"`
	Region             string      `json:"region"`
}

type PlayerGcoTokens struct {
	Class         string `json:"__class"`
	IdToken       string `json:"idToken"`
	UserInfoJwt   string `json:"userInfoJwt"`
	SummonerToken string `json:"summonerToken"`
}

type GameLobby struct {
	Class              string             `json:"__class"`
	PracticeGameConfig PracticeGameConfig `json:"practiceGameConfig"`
	SimpleInventoryJwt string             `json:"simpleInventoryJwt"`
	PlayerGcoTokens    PlayerGcoTokens    `json:"playerGcoTokens"`
}

type CustomLobbyResponse struct {
	Body struct {
		GameTypeConfigId float64 `json:"gameTypeConfigId"` //возвращает на подобии 12345.0 вместо 12345 (WTF RITO)
		Id               float64 `json:"id"`               //возвращает на подобии 12345.0 вместо 12345 (WTF RITO)
	} `json:"body"`
}

type LeagueSessionToken string

type RSOInventoryJWT string
