package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nerdwerx/daggerbot/config"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	for _, guild := range r.Guilds {
		gid := guild.ID

		guildData, err := s.Guild(gid)
		if err != nil {
			log.Printf("Error fetching guild data for %s: %v", gid, err)
			continue
		}

		config.Guilds[gid] = config.Guild{
			Name:  guildData.Name,
			Roles: guildData.Roles,
		}

		log.Printf("Connected to guild: %q (ID: %s)", config.Guilds[gid].Name, gid)

		if config.Debug {
			log.Printf("[DEBUG] Fetched %s", config.Guilds[gid])
		}
	}

	log.Printf("Bot is ready! Connected to %d servers", len(config.Guilds))
}
