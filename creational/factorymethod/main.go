package main

import "fmt"

func main() {
	content1, err := produceAudioContent("music")
	if err != nil {
		fmt.Println(err)
		return
	}

	content2, err := produceAudioContent("podcast")
	if err != nil {
		fmt.Println(err)
		return
	}

	content1.Play()
	content1.Transcribe()

	content2.Play()
	content2.Transcribe()
}
