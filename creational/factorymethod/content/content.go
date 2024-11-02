package content

type Type string

const (
	TypeMusic   Type = "music"
	TypePodcast Type = "podcast"
)

type AudioContent interface {
	Play()
	Transcribe()
}
