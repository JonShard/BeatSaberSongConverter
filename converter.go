package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ################ Old Format: ###########################################

/* xyonico's explenation of his format from BeatSaberSongLoader on https://github.com/xyonico/BeatSaberSongLoader.
"songName" - Name of your song
"songSubName" - Text rendered in smaller letters next to song name. (ft. Artist)
"beatsPerMinute" - BPM of the song you are using
"previewStartTime" - How many seconds into the song the preview should start
"previewDuration" - Time in seconds the song will be previewed in selection screen
"coverImagePath" - Cover image name
"environmentName" - Game environment to be used
"songTimeOffset" - Time in seconds of how early a song should start. Negative numbers for starting the song later
"shuffle" - Time in number of beats how much a note should shift
"shufflePeriod" - Time in number of beats how often a note should shift. Don't ask me why this is a feature, I don't know
"oneSaber" - true or false if it should appear in the one saber list
"difficultyLevels": [
	{
		"difficulty": This can only be set to Easy, Normal, Hard, Expert or ExpertPlus,
		"difficultyRank": Currently unused whole number for ranking difficulty,
		"jsonPath": The name of the json file for this specific difficulty
	}
*/

type DifficultyLevel struct {
	Difficulty     string  `json:"difficulty"`
	DifficultyRank int     `json:"difficultyRank"`
	AudioPath      string  `json:"audioPath"`
	JsonPath       string  `json:"jsonPath"`
	Offset         float32 `json:"offset"`
	OldOffset      float32 `json:"oldOffset"`
	ChromaToggle   string  `json:"chromaToggle"` // Why did the modding community make this a string, lol? Sounds like a binary value.
}

// InfoIn: The old structure of song info files.
type InfoIn struct {
	SongName         string            `json:"songName"`
	SongSubName      string            `json:"songSubName"`
	SongAuthorName   string            `json:"authorName"`
	BeatsPerMinute   float32           `json:"beatsPerMinute"`
	PreviewStartTime float32           `json:"previewStartTime"`
	PreviewDuration  float32           `json:"previewDuration"`
	CoverImagePath   string            `json:"coverImagePath"`
	EnvironmentName  string            `json:"environmentName"`
	OneSaber         bool              `json:"oneSaber"`
	DifficultyLevels []DifficultyLevel `json:"difficultyLevels"` // Array of DiffucultyLevel stucts, making a nested json object when encoded.
	// There are one of these for each difficulty in the level.
}

// ################ New Format: ###########################################

type DifficultyBeatmap struct {
	Difficulty         string  `json:"_difficulty"`
	DifficultyRank     int     `json:"_difficultyRank"`
	BeatMapFile        string  `json:"_beatmapFilename"`
	NoteJumpMovement   float32 `json:"_noteJumpMovementSpeed"`
	NoteJumpBeatOffset int     `json:"_noteJumpStartBeatOffset"`
}

type DifficultyBeatmapSet struct {
	BeatmapCharacteristicName string              `json:"_beatmapCharacteristicName"`
	DifficultyBeatMaps        []DifficultyBeatmap `json:"_difficultyBeatmaps"`
}

//InfoOut: The new structure of song info files. We will be converting to this one.
type InfoOut struct {
	Version               string                 `json:"_version"`
	SongName              string                 `json:"_songName"`
	SongSubName           string                 `json:"_songSubName"`
	SongAuthorName        string                 `json:"_songAuthorName"`
	LevelAuthorName       string                 `json:"_levelAuthorName"`
	BeatsPerMinute        float32                `json:"_beatsPerMinute"`
	SongTimeOffset        float32                `json:"_songTimeOffset"`
	Shuffle               float32                `json:"_shuffle"`
	ShufflePeriod         float32                `json:"_shufflePeriod"`
	PreviewStartTime      float32                `json:"_previewStartTime"`
	PreviewDuration       float32                `json:"_previewDuration"`
	SongFilename          string                 `json:"_songFilename"`
	CoverImageFileName    string                 `json:"_coverImageFilename"`
	EnvironmentName       string                 `json:"_environmentName"`
	DifficultyBeatmapSets []DifficultyBeatmapSet `json:"_difficultyBeatmapSets"`
}

func main() {

	// Get path to the song's info.json file:
	arg := os.Args[1]
	var infoFileName string
	if len(arg) != 0 {
		if strings.Contains(arg, "/info.json") {
			infoFileName = arg
		} else {
			infoFileName = arg + "/info.json"
		}
	} else {
		infoFileName = "info.json"
	}

	var infoIn InfoIn   // InfoIn struct for reading in the data from file.
	var infoOut InfoOut // InfoOut struct for writing the info file with the new format.

	// Open the info file:
	infoInFile, err := os.Open(infoFileName)
	if err != nil {
		fmt.Printf("opening info file: %s\n %s\n", infoFileName, err.Error())
		return
	}

	// Read the old info file.
	jsonParser := json.NewDecoder(infoInFile)
	if err = jsonParser.Decode(&infoIn); err != nil {
		fmt.Print("parsing info file\n", err.Error())
	}
	// Is the info file invalid because it doesen't have any difficulties?
	if len(infoIn.DifficultyLevels) == 0 {
		fmt.Printf("error parsing info file. There are no difficulty structures in the info.json file.\nThey are used to define presets of the song, each one using a separate Easy.json to ExpertPlus.json\n")
		return
	}

	// Assign values in the OUT struct with values from the IN struct.
	infoOut.Version = "2.0.0" // Magic number only 2 works. Maybe version of beat sabers native song loader? Or version of the format. :thinking:
	infoOut.SongName = infoIn.SongName
	infoOut.SongSubName = infoIn.SongSubName
	infoOut.SongAuthorName = infoIn.SongAuthorName
	infoOut.LevelAuthorName = "Unknown" // Although many level creators write their name in the SongSubName field. There is no guarantee for this.
	infoOut.BeatsPerMinute = infoIn.BeatsPerMinute
	infoOut.SongTimeOffset = 0 // Not sure what this offset is.
	infoOut.Shuffle = 0.0
	infoOut.ShufflePeriod = 0.5 // Not sure what this is. 0.5 seems to be the default value when making a level with the in-game editor.
	infoOut.PreviewStartTime = infoIn.PreviewStartTime
	infoOut.PreviewDuration = infoIn.PreviewDuration
	infoOut.SongFilename = infoIn.DifficultyLevels[0].AudioPath // Beat saber's format doesn't support several audio files for different difficulties :( oh well, taking the first one.
	infoOut.CoverImageFileName = infoIn.CoverImagePath
	infoOut.EnvironmentName = infoIn.EnvironmentName

	var difficultySet DifficultyBeatmapSet
	difficultySet.BeatmapCharacteristicName = "Standard" // Standard, One Saber, No Arrows. As far as I see, the old format doesn't support separate modes for difficulties, meaning that the entire song is one hand or not.
	for i := 0; i < len(infoIn.DifficultyLevels); i++ {

		var beatmap DifficultyBeatmap
		beatmap.Difficulty = infoIn.DifficultyLevels[i].Difficulty
		beatmap.DifficultyRank = infoIn.DifficultyLevels[i].DifficultyRank
		beatmap.BeatMapFile = infoIn.DifficultyLevels[i].JsonPath
		beatmap.NoteJumpMovement = 10.0 // Not sure what this is. But 10 seems to be a default value in the beat saber level editor.
		beatmap.NoteJumpBeatOffset = 0
		difficultySet.DifficultyBeatMaps = append(difficultySet.DifficultyBeatMaps, beatmap)
	}
	infoOut.DifficultyBeatmapSets = append(infoOut.DifficultyBeatmapSets, difficultySet)

	// We're done making the info file, time to write it to disk. Make a new Info.dat file.
	infoOutFile, err := os.Create(strings.Replace(infoFileName, "info.json", "Info.dat", 1))
	if err != nil {
		fmt.Print("Unable to create Info.dat file.\n", err.Error())
	}

	// Encode the OUT json structure. Resulting in a
	jsonEncoder := json.NewEncoder(infoOutFile)
	err = jsonEncoder.Encode(infoOut)
	if err != nil {
		fmt.Print("Unable to encode json structure\n", err.Error())
	}

	fmt.Printf("Successfully converted %s's 'info.json' into an 'Info.dat'\n", infoOut.SongName)

	// Close the files:
	if infoInFile != nil { // Close read Info file.
		err = infoInFile.Close()
	}
	if infoOutFile != nil {
		err = infoOutFile.Close()
	}

	return
}
