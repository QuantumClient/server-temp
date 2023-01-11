package services

import (
	"github.com/bwmarrin/discordgo"
	"quantumclient.org/backend/v2/models"
	"strings"
	"time"
)

type DiscordService struct {
	Conf *models.Config
	Discord *discordgo.Session
}

func NewDiscordService(conf *models.Config) *DiscordService {
	return &DiscordService{
		Conf: conf,
	}
}

func (s DiscordService) Init() {
	ds, err := discordgo.New("Bot " + s.Conf.Discord.Token)
	if err != nil {
		println("error creating Discord session,", err)
		return
	}

	s.Discord = ds
	ds.AddHandler(
		func(s *discordgo.Session, r *discordgo.Ready) {
			println("Discord ready - logged in as " + r.User.String())
		},
	)
	ds.AddHandler(
		func(se *discordgo.Session, m *discordgo.MessageCreate) {
			if strings.Contains(m.Content, "steam") && strings.Contains(m.Content, "free") {
				t, err := m.Member.JoinedAt.Parse()
				if err != nil {
					println("error parsing time", err)
					return
				}
				if time.Since(t).Hours() < 36 {
					se.GuildMemberRoleAdd(m.GuildID, m.Author.ID, s.Conf.Discord.MuteRole)
					se.ChannelMessageSend(m.ChannelID, "You have been muted for spam")
				}
			}
		},
	)
	s.Discord.Open()

}

func (s DiscordService) End() {
	s.Discord.Close()
}

