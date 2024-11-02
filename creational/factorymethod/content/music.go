package content

import "fmt"

type Music struct {
	category      string
	transcription bool
}

func NewMusic() AudioContent {
	return &Music{
		category:      "music",
		transcription: false,
	}
}

func (m *Music) Play() {
	fmt.Printf("play %s...\n", m.category)
}

func (p *Music) Transcribe() {
	if !p.transcription {
		fmt.Printf("%s cannot be transcribed\n", p.category)
		return
	}

	fmt.Printf("transcribing %s...\n", p.category)
}
