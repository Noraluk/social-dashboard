package utils

import (
	"image/color"
	"image/png"
	"os"
	"sort"

	"github.com/psykhi/wordclouds"
)

var DefaultColors = []color.RGBA{
	{0x1b, 0x1b, 0x1b, 0xff},
	{0x48, 0x48, 0x4B, 0xff},
	{0x59, 0x3a, 0xee, 0xff},
	{0x65, 0xCD, 0xFA, 0xff},
	{0x70, 0xD6, 0xBF, 0xff},
}

type Conf struct {
	FontMaxSize     int          `json:"font_max_size"`
	FontMinSize     int          `json:"font_min_size"`
	RandomPlacement bool         `json:"random_placement"`
	FontFile        string       `json:"font_file"`
	Colors          []color.RGBA `json:"colors"`
	Width           int          `json:"width"`
	Height          int          `json:"height"`
}

var DefaultConf = Conf{
	RandomPlacement: false,
	FontFile:        "./assets/fonts/Sriracha-Regular.ttf",
	Colors:          DefaultColors,
	Width:           3072,
	Height:          3072,
}

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func GererateWordClouds(word map[string]int, fileName string, maxWord, maxFontSize, minFontSize int) error {
	var words []WordCount
	for w, count := range word {
		words = append(words, WordCount{Word: w, Count: count})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})

	wordCount := make(map[string]int)
	for i := 0; i < maxWord; i++ {
		wordCount[words[i].Word] = words[i].Count
	}

	colors := make([]color.Color, 0)
	for _, c := range DefaultConf.Colors {
		colors = append(colors, c)
	}

	w := wordclouds.NewWordcloud(wordCount,
		wordclouds.FontFile(DefaultConf.FontFile),
		wordclouds.FontMaxSize(maxFontSize),
		wordclouds.FontMinSize(minFontSize),
		wordclouds.Colors(colors),
		wordclouds.Height(DefaultConf.Height),
		wordclouds.Width(DefaultConf.Width),
		wordclouds.RandomPlacement(DefaultConf.RandomPlacement),
	)
	img := w.Draw()

	outputFile, err := os.Create(fileName)
	defer outputFile.Close()
	if err != nil {
		return err
	}

	err = png.Encode(outputFile, img)
	if err != nil {
		return err
	}
	return nil
}
