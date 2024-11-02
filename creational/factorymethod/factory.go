package main

import (
	"errors"
	"factorymethod/content"
)

func produceAudioContent(contentType content.Type) (content.AudioContent, error) {
	if contentType == content.TypeMusic {
		return content.NewMusic(), nil
	}

	if contentType == content.TypePodcast {
		return content.NewPodcast(), nil
	}

	return nil, errors.New("invalid content type")
}
