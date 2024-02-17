package lcu

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ImOlli/go-lcu/lcu"
	"io"
	"net/http"
	"net/url"
	"time"
)

const URL = "https://127.0.0.1:%s"

func NewLCU() (*LCU, error) {
	info, err := lcu.FindLCUConnectInfo()
	if err != nil {
		return nil, err
	}

	destURL := fmt.Sprintf(URL, info.Port)

	netTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: netTransport,
	}

	return &LCU{
		DestURL: destURL,
		Info:    info,
		Client:  client,
	}, nil
}

func (lcu *LCU) GetRSOIdToken() (*RSOIdToken, error) {
	var token = &RSOIdToken{}

	response, err := lcu.MakeGetRequest(RSOIdTokenMethod)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(response, &token)

	return token, nil
}

func (lcu *LCU) GetLeagueSessionToken() (*LeagueSessionToken, error) {
	response, err := lcu.MakeGetRequest(LeagueSessionTokenMethod)
	if err != nil {
		return nil, err
	}

	return NewLeagueSessionToken(string(response)), nil
}

func (lcu *LCU) GetRSOInventoryJWT() (*RSOInventoryJWT, error) {
	response, err := lcu.MakeGetRequest(RSOInventoryJWTMethod)
	if err != nil {
		return nil, err
	}

	return NewRSOInventoryJWT(string(response)), nil
}

func (lcu *LCU) GetRSOAccessToken() (*RSOAccessToken, error) {
	var accessToken = &RSOAccessToken{}

	response, err := lcu.MakeGetRequest(RSOAccessTokenMethod)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(response, &accessToken)

	return accessToken, nil
}

func (lcu *LCU) GetGameVersion() (string, error) {
	response, err := lcu.MakeGetRequest(GameVersionMethod)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func (lcu *LCU) GetChampSelectSession() (*ChampSelectSession, error) {
	var session = &ChampSelectSession{}

	response, err := lcu.MakeGetRequest(ChampSelectSessionMethod)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(response, &session)

	return session, nil
}

func (lcu *LCU) GetMySelection() (*LobbySummoner, error) {
	var summoner = &LobbySummoner{}

	response, err := lcu.MakeGetRequest(MySelectionMethod)
	if err != nil {
		return nil, err
	}

	_ = json.Unmarshal(response, &summoner)

	return summoner, nil
}

func (lcu *LCU) QuitCustomLobby() ([]byte, error) {
	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "quitGame")
	params.Add("args", "[]")

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) SendCustomLobby(lobby *GameLobby) (*CustomLobbyResponse, error) {
	var lobbyResponse CustomLobbyResponse

	jsonLobby, err := json.Marshal([]GameLobby{*lobby})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "createPracticeGameV4")
	params.Add("args", string(jsonLobby))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &lobbyResponse)
	if err != nil {
		return nil, err
	}

	return &lobbyResponse, nil
}

func (lcu *LCU) StartChampionSelection(id, gameTypeConfigId int) ([]byte, error) {
	jsonArgs, err := json.Marshal([]int{id, gameTypeConfigId})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "startChampionSelection")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) SetClientReceivedGameMessage(id int) ([]byte, error) {
	jsonArgs, err := json.Marshal([]any{id, "CHAMP_SELECT_CLIENT"})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "setClientReceivedGameMessage")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) SelectSpells(spell1, spell2 int) ([]byte, error) {
	jsonArgs, err := json.Marshal([]int{spell1, spell2})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "selectSpells")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) SelectChampionV2(championId, skinId int) ([]byte, error) {
	jsonArgs, err := json.Marshal([]int{championId, skinId})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "selectChampionV2")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) ChampionSelectCompleted() ([]byte, error) {
	jsonArgs, err := json.Marshal([]int{})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "championSelectCompleted")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) SetClientReceivedMaestroMessage(id int) ([]byte, error) {
	jsonArgs, err := json.Marshal([]any{id, "GameClientConnectedToServer"})
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("destination", "gameService")
	params.Add("method", "setClientReceivedMaestroMessage")
	params.Add("args", string(jsonArgs))

	response, err := lcu.MakePostRequest(SessionInvokeMethod, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (lcu *LCU) MakeGetRequest(method string) ([]byte, error) {
	endpoint := lcu.DestURL + method

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header = makeLCUHeader(lcu.Info.AuthToken)

	response, err := lcu.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	return responseBody, nil
}

func (lcu *LCU) MakePostRequest(method string, params url.Values) ([]byte, error) {
	endpoint := lcu.DestURL + method

	request, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header = makeLCUHeader(lcu.Info.AuthToken)
	request.URL.RawQuery = params.Encode()

	response, err := lcu.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, _ := io.ReadAll(response.Body)
	return responseBody, nil
}
