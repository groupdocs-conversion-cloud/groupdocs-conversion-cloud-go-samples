package info

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
)

func GetLicenseConsumption() {

	response, _, err := config.Client.LicenseApi.GetConsumptionCredit(config.Ctx)

	if err != nil {
		fmt.Printf("GetLicenseConsumption error: %v\n", err)
		return
	}

	fmt.Printf("Credits: %v\n", response.Credit)
	fmt.Printf("Quantity: %v\n", response.Quantity)
}
