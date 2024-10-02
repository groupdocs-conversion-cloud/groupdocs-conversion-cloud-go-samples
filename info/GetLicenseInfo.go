package info

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
)

func GetLicenseInfo() {

	response, _, err := config.Client.LicenseApi.GetLicenseInfo(config.Ctx)

	if err != nil {
		fmt.Printf("GetLicenseInfo error: %v\n", err)
		return
	}

	fmt.Printf("IsLicensed: %v\n", response.IsLicensed)
}
