package foot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type FootCommand struct {
	Receiver *Receiver
	payload  *FootCommandPayload
}

type FootCommandPayload struct {
	Answer  []*discordgo.MessageEmbed
	session *discordgo.Session
	Message *discordgo.MessageCreate
}

type LastCompetitions struct {
	Count   uint     `json:"count"`
	Matches []Matche `json:"matches"`
}

type Matche struct {
	HomeTeam Team `json:"homeTeam"`
	AwayTeam Team `json:"awayTeam"`
}

type Team struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func GetLastCompetitionsCommand(s *discordgo.Session, m *discordgo.MessageCreate) *FootCommand {
	req, err := http.NewRequest("GET", "https://api.football-data.org/v2/competitions/FL1/matches?matchday=10", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Set("X-Auth-Token", os.Getenv("API_KEY"))

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	var LastC LastCompetitions
	errJson := json.Unmarshal(body, &LastC)
	if errJson != nil {
		log.Println(errJson)
	}
	var MessagesEmbbed []*discordgo.MessageEmbed
	for i := 0; i < len(LastC.Matches); i++ {
		log.Print(LastC.Matches[i].AwayTeam.Name)
		log.Print(LastC.Matches[i].HomeTeam.Name)
		log.Print("--------------")
		tempEmbbed := discordgo.MessageEmbed{Title: "Super match", Description: fmt.Sprintf("%s VS %s", LastC.Matches[i].AwayTeam.Name, LastC.Matches[i].HomeTeam.Name)}

		MessagesEmbbed = append(MessagesEmbbed, &tempEmbbed)
	}

	return &FootCommand{
		Receiver: &Receiver{},
		payload: &FootCommandPayload{
			Answer:  MessagesEmbbed,
			session: s,
			Message: m,
		},
	}
}

func (c *FootCommand) Execute() {
	c.Receiver.Answer(c.Payload())
}

func (c *FootCommand) Payload() *FootCommandPayload {
	return c.payload
}

func (p *FootCommandPayload) Session() *discordgo.Session {
	return p.session
}
