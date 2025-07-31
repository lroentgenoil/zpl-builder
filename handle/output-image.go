package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
	"github.com/lroentgenoil/zpl-builder/elements"
)

func ImageOutput(res []zebrashElements.LabelInfo, params elements.FormattedParams, opts drawers.DrawerOptions) error {
	drawer := zebrash.NewDrawer()

	var result interface{}
	switch params.Output {
	case "file":
		result = []string{}
	default:
		result = [][]byte{}
	}

	for idx, label := range res {
		buf := &bytes.Buffer{}
		err := drawer.DrawLabelAsPng(label, buf, opts)
		if err != nil {
			return fmt.Errorf("error generando imagen: %w", err)
		}

		switch params.Output {
		case "file":
			filepath := fmt.Sprintf(params.UrlOutput+"imagen%05d.png", idx)
			err := os.WriteFile(filepath, buf.Bytes(), 0644)
			if err != nil {
				return fmt.Errorf("error generando imagen: %w", err)
			}

			result = append(result.([]string), filepath)
		default:
			result = append(result.([][]byte), buf.Bytes())
		}
	}

	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error generando Json: %w", err)
	}

	_, err = os.Stdout.Write(data)
	return err
}
