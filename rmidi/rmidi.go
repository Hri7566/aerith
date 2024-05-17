package rmidi

import "gitlab.com/gomidi/midi/v2/smf"

var Data *smf.SMF

func LoadMIDI(path string) (err error) {
	mf, err := smf.ReadFile(path)

	if err != nil {
		return err
	}

	Data = mf

	return nil
}

func UnloadMIDI() {
	Data = nil
}
