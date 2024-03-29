package telegram

import (
	tgbotapi "github.com/mohammadkarimi23/telegram-bot-api/v5"
	"github.com/spotify-bot/server/pkg/spotify"
	"github.com/spotify-bot/telegram/internal/config"
)

func getAuthMessage(userID string) tgbotapi.InlineKeyboardMarkup {

	link := config.AppConfig.APIServerAddress + "/auth/telegram?user_id=" + userID //FIXME dev config
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Login to Spotify", link),
		),
	)
}

func getTrackQueryResult(id string, track *spotify.Track) tgbotapi.InlineQueryResultArticle {
	trackLink := spotify.OpenSpotifyTrackEndpoint + track.ID
	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Play Now", "1#"+track.ID),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Add to Queue", "2#"+track.ID),
		),
	)
	return tgbotapi.InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: track.Name,
		InputMessageContent: tgbotapi.InputTextMessageContent{
			Text: trackLink,
		},
		ReplyMarkup: &replyMarkup,
		ThumbURL:    track.Album.Images[1].URL, //FIXME use smallest image
		ThumbWidth:  track.Album.Images[1].Width,
		ThumbHeight: track.Album.Images[1].Height,
		Description: track.Album.Artists[0].Name, //FIXME
	}
}
