package commands

import (
	"encoding/json"
	"go-discord-bot/utils"
	"io"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	apiURL    = "https://i.instagram.com/api/v1/users/web_profile_info/?username="
	userAgent = "iphone_ua"
	igAppID   = "936619743392459"
)

type userData struct {
	Data struct {
		User struct {
			Username     string `json:"username"`
			FullName     string `json:"full_name"`
			HdProfileUrl string `json:"profile_pic_url_hd"`
			ProfileUrl   string `json:"profile_pic_url"`
			Bio          string `json:"biography"`
			FollowedBy   struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
			Follows struct {
				Count int `json:"count"`
			} `json:"edge_follow"`
		} `json:"user"`
	} `json:"data"`
}

func fetchInstagramData(username string) (*userData, error) {
	url := apiURL + username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("x-ig-app-id", igAppID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data userData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func Insta(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := strings.Split(m.Content, " ")
	if len(args) <= 1 {
		errorEmbed := &discordgo.MessageEmbed{
			Description: "Please provide an Instagram username",
			Color:       0xC20C00,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, errorEmbed)
		return
	}

	username := strings.Join(args[1:], " ")
	data, err := fetchInstagramData(username)
	if err != nil {
		errorEmbed := &discordgo.MessageEmbed{
			Description: "An error occurred while fetching Instagram data",
			Color:       0xC20C00,
		}
		s.ChannelMessageSendEmbed(m.ChannelID, errorEmbed)
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       data.Data.User.FullName,
		Description: data.Data.User.Bio,
		URL:         "https://www.instagram.com/" + data.Data.User.Username,
		Color:       0x00B6C2,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Followers",
				Value:  utils.FormatNumber(data.Data.User.FollowedBy.Count),
				Inline: true,
			},
			{
				Name:   "Following",
				Value:  utils.FormatNumber(data.Data.User.Follows.Count),
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: data.Data.User.ProfileUrl,
		},
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
