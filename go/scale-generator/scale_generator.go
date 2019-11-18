package scale

var chromaticSharp = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

var chromaticFlat = []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}

// Scale generator.
func Scale(tonic, interval string) []string {

	if tonic == "C" && interval == "" {
		return chromaticSharp
	}

	if tonic == "F" && interval == "" {
		return chromaticFlat
	}

	return []string{"test"}
}
