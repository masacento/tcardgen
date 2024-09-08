package png

import (
	"image/png"
	"os"

	"github.com/xyproto/palgen"
)

func Optimize(inpath, outpath string, colors int) error {
	inputFile, err := os.Open(inpath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	img, err := png.Decode(inputFile)
	if err != nil {
		return err
	}

	palette, err := palgen.Generate(img, colors)
	if err != nil {
		return err
	}

	convertedImg, err := palgen.ConvertCustom(img, palette)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return png.Encode(outputFile, convertedImg)
}
