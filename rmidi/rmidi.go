package rmidi

import "gitlab.com/gomidi/midi/v2/smf"

var MIDIEvents []smf.TrackEvent

func LoadMIDI(path string) {
	smf.ReadTracks(path).Do(
		func(te smf.TrackEvent) {
			MIDIEvents = append(MIDIEvents, te)
		},
	)
}

func UnloadMIDI() {
	MIDIEvents = []smf.TrackEvent{}
}
