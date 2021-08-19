package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jemurai/fkit/finding"

	log "github.com/sirupsen/logrus"
)

func PrintFindings(findings []finding.Finding) {
	fjson, _ := json.MarshalIndent(findings, "", " ")
	fmt.Printf("%s", fjson)
}

// BuildFindingsFromNPMLCFile read a json file of Findings and build an array
// of findings that can be used for further processing.
func BuildFindingsFromNPMLCFile(file string) []finding.Finding {
	var findings []finding.Finding
	var npmlcreportmap map[string]LicenseDetail

	rfile, err := os.Open(file)
	if err != nil {
		log.Error(err)
	}
	bytes, err := ioutil.ReadAll(rfile)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(bytes, &npmlcreportmap)
	if err != nil {
		log.Error(err)
	}
	log.Debugf("Unmarshalled NPM License Check JSON")
	num := 0
	for key, license := range npmlcreportmap {
		var refs []string
		var tags []string
		var cwes []string
		tags = append(tags, "npm license-checker")

		var source string = "LicenceChecker: " + license.Path
		var description string = license.Repository + " with license: " + license.Licenses
		var detail string = license.Licenses + " at " + license.Path + " from: " + license.URL + " by: " + license.Publisher + " with license at: " + license.LicenseFile
		name := key + " : " + license.Licenses
		hasher := sha256.New()
		hasher.Write([]byte(license.Path + name))
		fingerprint := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		finding := finding.Finding{
			Name:        name,
			Description: description,
			Detail:      detail,
			//Severity:    vuln.Severity,
			//Confidence:  vuln.Confidence,
			Fingerprint: fingerprint,
			Timestamp:   time.Now(),
			Source:      source,
			Location:    license.Path,
			Cvss:        0,
			References:  refs,
			Cwes:        cwes,
			Tags:        tags,
		}
		if name != "" {
			num++
			findings = append(findings, finding)
		}
	}
	log.Debugf("NPM License Check To Off Processed %v vulns", num)
	return findings
}
