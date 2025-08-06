package handle

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
	"github.com/lroentgenoil/zpl-builder/elements"
	"github.com/lroentgenoil/zpl-builder/jobs"
)

func ImageOutput(res []zebrashElements.LabelInfo, params elements.FormattedParams, opts drawers.DrawerOptions) error {
	lote := params.Chunk
	var result interface{}
	switch params.Output {
	case "file":
		result = []string{}
	default:
		result = [][]byte{}
	}

	for start := 0; start < len(res); start += lote {
		end := start + lote
		if end > len(res) {
			end = len(res)
		}

		currentLote := res[start:end]

		outputBuffers, err := jobs.GeneratePNGsParallel(currentLote, opts)
		if err != nil {
			return err
		}

		for idx, buf := range outputBuffers {
			if buf == nil {
				continue
			}

			index := start + idx

			switch params.Output {
			case "file":
				filepath := fmt.Sprintf(params.UrlOutput+"imagen%05d.png", index)
				err := os.WriteFile(filepath, buf.Bytes(), 0644)
				if err != nil {
					return fmt.Errorf("error generando imagen: %w", err)
				}

				result = append(result.([]string), filepath)
			default:
				result = append(result.([][]byte), buf.Bytes())
			}
		}
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error generando Json: %w", err)
	}

	_, err = os.Stdout.Write(data)
	return err
}
