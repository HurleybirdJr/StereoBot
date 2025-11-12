package bot

import (
	"slices"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) DoesMemberHaveAnyRoles(memberRoles []string, checkRoles []string) bool {
	if len(checkRoles) == 0 {
		return true
	}
	return slices.ContainsFunc(memberRoles, func(e string) bool {
		for _, role := range checkRoles {
			if b.Config.Roles[role] == e {
				return true
			}
		}
		return false
	})
}

func (b *Bot) IsMemberDriver(member *discordgo.Member) bool {
	modRoleId := b.Config.Roles["Driver"]
	return slices.Contains(member.Roles, modRoleId)
}
