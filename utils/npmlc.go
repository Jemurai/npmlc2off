package utils

// LicenseDetail captures a license.
type LicenseDetail struct {
	URL         string `json:"url"`
	Path        string `json:"path"`
	Licenses    string `json:"licenses"`
	Publisher   string `json:"publisher"`
	Repository  string `json:"repository"`
	Email       string `json:"email"`
	LicenseFile string `json:"licenseFile"`
}
