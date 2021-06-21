package dot2pdf

import (
	"os/exec"
	"powershellDeal/common/logger"
)

func DotToPDF(dotpath, pdfpath string) {
	out, err := exec.Command("dot", "-Tpdf", dotpath, "-o", pdfpath).Output()
	if err != nil {
		logger.Logger.Warn(err)
	} else {
		if len(out) != 0 {
			logger.Logger.Info(out)
		}

	}
}
