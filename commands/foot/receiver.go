package foot

type Receiver struct{}

func (r *Receiver) Answer(p *FootCommandPayload) {
	for i := 0; i < len(p.Answer); i++ {
		p.Session().ChannelMessageSendEmbed(p.Message.ChannelID, p.Answer[i])
	}
}
