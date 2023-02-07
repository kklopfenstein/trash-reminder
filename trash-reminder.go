package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Resp struct {
	Events []Event
}

type Event struct {
	Flags []Flag
	Day   string
}

type Flag struct {
	Id        int
	Name      string
	EventType string `json:"event_type"`
	Subject   string `json:"subject"`
	Html      string `json:"html_message"`
	Icon      string `json:"icon"`
}

func main() {
	place := flag.String("place", "A0DDD532-8B87-11E7-96F5-A314FC2257CE", "recollect api place, e.g. ")
	service := flag.String("service", "237", "recollect api service")
	discordUserId := flag.String("discordUserId", "", "discord user id")
	discordToken := flag.String("discordToken", "", "discord token")
	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("https://api.recollect.net/api/places/%s/services/%s/events", *place, *service))
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var eventsResp Resp
	err = json.Unmarshal(data, &eventsResp)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d events", len(eventsResp.Events))

	message := make([]string, 0)

	for _, event := range eventsResp.Events {
		log.Printf("%d flags", len(event.Flags))
		for _, flag := range event.Flags {
			log.Printf("%s\t%s", event.Day, flag.Name)
			t, err := time.Parse("2006-01-02", event.Day)
			log.Printf("date is %s", t)
			if err != nil {
				log.Fatalf("Could not parse date %s", event.Day)
			}
			dur := t.Sub(time.Now())
			log.Printf("Diff in hours is %f", dur.Hours())
			if dur.Hours() < 24 && dur.Hours() > 0 {
				message = append(message, fmt.Sprintf("%s collection today!", toReadableName(flag.Name)))
			}
		}
	}

	discord, err := discordgo.New("Bot " + *discordToken)
	defer discord.Close()
	if err != nil {
		log.Fatal(err)
	}
	channel, err := discord.UserChannelCreate(*discordUserId)
	if err != nil {
		log.Fatal(err)
	}
	_, err = discord.ChannelMessageSend(channel.ID, strings.Join(message, "\n"))
}

func toReadableName(t string) string {
	switch t {
	case "garbage":
		return "Garbage"
	case "recycling":
		return "Recycling"
	case "yardtrimmings":
		return "Yard waste"
	case "notrashcollection":
		return "No trash"
	}
	return t
}
