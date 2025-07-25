package handle

import (
	"bytes"
	"fmt"
	"os"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	zebrashElements "github.com/lroentgenoil/zebrashMod/elements"
)

func ImageOutput(label []zebrashElements.LabelInfo, opts drawers.DrawerOptions) error {
	buf := &bytes.Buffer{}
	drawer := zebrash.NewDrawer()

	err := drawer.DrawLabelAsPng(label[0], buf, opts)
	if err != nil {
		return fmt.Errorf("error generando imagen: %w", err)
	}

	_, err = os.Stdout.Write(buf.Bytes())
	return err
}
