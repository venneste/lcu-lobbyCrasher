package main

import (
	lcu2 "lcu-lobbyCrasher/lcu"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	lcu, err := lcu2.NewLCU()
	if err != nil {
		log.Fatalf("Откройте клиент лиги")
		return
	}
	log.Printf("Клиент лиги запущен на %s (логин: riot, пароль: %s)\n", lcu.DestURL, lcu.Info.AuthToken)

	if !prepareForCustomLobby(lcu) {
		return
	}

	log.Println("Генерация кастом лобби...")
	customLobby, err := generateCustomLobby(lcu)
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке сгенерировать кастом лобби: %s", err.Error())
		return
	}

	log.Println("Отправка LCU кастом лобби...")
	lobby, err := lcu.SendCustomLobby(customLobby)
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке отправить кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.StartChampionSelection(int(lobby.Body.Id), int(lobby.Body.GameTypeConfigId))
	if err != nil {
		log.Fatalf("Произошла ошибка в кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.SetClientReceivedGameMessage(int(lobby.Body.Id))
	if err != nil {
		log.Fatalf("Произошла ошибка в кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.SelectSpells(32, 4)
	if err != nil {
		log.Fatalf("Произошла ошибка в кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.SelectChampionV2(1, 1000)
	if err != nil {
		log.Fatalf("Произошла ошибка в кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.ChampionSelectCompleted()
	if err != nil {
		log.Fatalf("Произошла ошибка в кастом лобби: %s", err.Error())
		return
	}

	log.Println("Подождите 15 секунд...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(15 * time.Second)
	}()

	wg.Wait()

	_, err = lcu.QuitCustomLobby()
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке выйти из кастом лобби: %s", err.Error())
		return
	}

	_, err = lcu.SetClientReceivedMaestroMessage(int(lobby.Body.Id))
	if err != nil {
		log.Fatalf("Произошла ошибка при отправке последнего запроса: %s", err.Error())
		return
	}

	log.Println("Лобби крашнулось!")
	os.Exit(0)
}

func prepareForCustomLobby(lcu *lcu2.LCU) bool {
	_, err := lcu.QuitCustomLobby()
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке выйти из кастом лобби: %s", err.Error())
		return false
	}

	session, err := lcu.GetChampSelectSession()
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке получить текущее лобби: %s", err.Error())
		return false
	}
	if session.IsCustomGame {
		log.Fatalf("Вы не можете крашнуть кастомное лобби")
		return false
	}

	mySelection, err := lcu.GetMySelection()
	if err != nil {
		log.Fatalf("Произошла ошибка при попытке получить вашу информацию в текущем лобби: %s", err.Error())
		return false
	}
	if mySelection.ChampionId == 0 {
		log.Fatalf("Вам необходимо подтвердить выбор чемпиона перед крашом лобби")
		return false
	}

	remaining := int(session.Timer.InternalNowInEpochMs-time.Now().UTC().UnixMilli()) + session.Timer.AdjustedTimeLeftInPhase
	if (session.Timer.Phase == "FINALIZATION") && (remaining < 11000) {
		log.Fatalf("Необходимо как минимум 11 секунд для краша лобби")
		return false
	}

	return true
}

func generateCustomLobby(lcu *lcu2.LCU) (*lcu2.GameLobby, error) {
	gameVersion, err := lcu.GetGameVersion()
	if err != nil {
		return nil, err
	}
	inventoryJwt, err := lcu.GetRSOInventoryJWT()
	if err != nil {
		return nil, err
	}
	idToken, err := lcu.GetRSOIdToken()
	if err != nil {
		return nil, err
	}
	accessToken, err := lcu.GetRSOAccessToken()
	if err != nil {
		return nil, err
	}
	sessionToken, err := lcu.GetLeagueSessionToken()
	if err != nil {
		return nil, err
	}

	return &lcu2.GameLobby{
		Class: "com.riotgames.platform.game.lcds.dto.CreatePracticeGameRequestDto",
		PracticeGameConfig: lcu2.PracticeGameConfig{
			Class:              "com.riotgames.platform.game.PracticeGameConfig",
			AllowSpectators:    "NONE",
			GameMap:            *lcu2.MakeGameMap(),
			GameMode:           "CLASSIC",
			GameMutators:       []int{},
			GameName:           "Summoners Rift 5v5",
			GamePassword:       "",
			GameTypeConfig:     1,
			GameVersion:        lcu2.TrimQuotes(gameVersion),
			MaxNumPlayers:      10,
			PassbackDataPacket: nil,
			PassbackUrl:        nil,
			Region:             "",
		},
		SimpleInventoryJwt: inventoryJwt.String(),
		PlayerGcoTokens: lcu2.PlayerGcoTokens{
			Class:         "com.riotgames.platform.util.tokens.PlayerGcoTokens",
			IdToken:       idToken.Token,
			UserInfoJwt:   accessToken.Token,
			SummonerToken: sessionToken.String(),
		},
	}, nil
}
