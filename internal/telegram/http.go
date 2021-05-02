package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spotify-bot/server/pkg/spotify"
	"github.com/spotify-bot/telegram/internal/config"
)

func getRecentlyPlayed(userID string) (track *spotify.Track, err error) {
	track, err = getCurrentlyPlayingSong(userID)
	if err != nil {
		track, err = getRecentlyPlayedSong(userID)
	}
	return
}

func getRecentlyPlayedSong(userID string) (*spotify.Track, error) {

	path := spotify.RecentlyPlayedEndpoint + "?limit=1"

	resp, err := sendRequest("GET", path, userID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response spotify.RecentlyPlayedResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	if len(response.Items) < 1 {
		return nil, fmt.Errorf("Empty track")
	}
	return &response.Items[0].Track, nil
}

func getCurrentlyPlayingSong(userID string) (*spotify.Track, error) {

	path := spotify.CurrentlyPlayingEndpoint
	resp, err := sendRequest("GET", path, userID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, spotify.CallbackError{
			Endpoint: spotify.CurrentlyPlayingEndpoint,
			Code:     resp.StatusCode,
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response spotify.CurrentlyPlayingResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response.Track, nil
}

func addSongToQueue(userID string, songURI string) error {
	resp, err := sendRequest("POST", spotify.AddToQueueEndpoint+"?uri="+songURI, userID, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("request error code: %d\n", resp.StatusCode)
	}
	return nil
}

func playSong(userID string, songURI string) error {

	var jsonStr = []byte(`{"uris":["` + songURI + `"]}`)
	resp, err := sendRequest("PUT", spotify.PlaySongEndpoint, userID, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("request error code: %d", resp.StatusCode)
	}
	return nil
}

func sendRequest(method string, path string, userID string, body io.Reader) (*http.Response, error) {

	client := &http.Client{}
	url := config.AppConfig.APIServerAddress + "/spotify/telegram/" + userID + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
