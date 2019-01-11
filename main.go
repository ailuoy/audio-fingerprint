package main

import (
	"fmt"
	"github.com/audio-fingerprint/fingerprint"
)

func main() {
	first, _ := fingerprint.GetFingerPrintInfo("first.mp3")
	second, _ := fingerprint.GetFingerPrintInfo("second.mp3")
	similarity, _, _ := fingerprint.Correlate(first.Fingerprint, second.Fingerprint)
	fmt.Println(similarity)
}
