package content

import "fmt"

type Podcast struct {
	category      string
	transcription bool
}

func NewPodcast() AudioContent {
	return &Podcast{
		category:      "podcast",
		transcription: true,
	}
}

func (p *Podcast) Play() {
	fmt.Printf("play %s...\n", p.category)
}

func (p *Podcast) Transcribe() {
	if !p.transcription {
		fmt.Printf("%s cannot be transcribed\n", p.category)
		return
	}

	fmt.Printf("transcribing %s...\n", p.category)
}
