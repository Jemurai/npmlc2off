package cmd

import (
	utils "github.com/jemurai/npmlc2off/utils"
)

func Convert(file string) {
	findings := utils.BuildFindingsFromNPMLCFile(file)
	utils.PrintFindings(findings)
}
