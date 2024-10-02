//// ***********************************************************
////          GroupDocs.Conversion Cloud API Examples
//// ***********************************************************

package main

import (
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/common"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/convert"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/convert/async"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/info"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/cad"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/csv"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/email"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/html"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/note"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/pdf"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/presentation"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/spreadsheet"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/txt"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/loadOptions/wordprocessing"
)

func main() {

	// TODO: Get your ClientId and ClientSecret at https://dashboard.groupdocs.cloud (free registration is required).
	config.InitAndUpload("Your_ClientId", "Your_ClientSecret")

	// Info API Examples
	info.GetDocumentInformation()
	info.GetSupportedConversions()

	// License API Examples (only for self-hosted API)
	// info.GetLicenseConsumption()
	// info.GetLicenseInfo()

	// Document conversion examples with common options
	common.AddWatermark()
	common.ConvertConsecutivePages()
	common.ConvertSpecificPages()
	common.ConvertUsingCustomFont()

	// Document async conversion examples
	async.ConvertToPdfAsync()
	async.ConvertToPdfDirectAsync()

	// Document conversion examples with conversion options
	convert.ConvertToHtml()
	convert.ConvertToImage()
	convert.ConvertToPdf()
	convert.ConvertToPdfDirect()
	convert.ConvertToPdfDirectOptions()
	convert.ConvertToPdfResponseBody()
	convert.ConvertToPresentation()
	convert.ConvertToSpreadsheet()
	convert.ConvertToWordProcessing()

	// Document conversion examples with load options
	cad.ConvertCadAndSpecifyLoadOptions()
	csv.ConvertCsvByConvertingDateTimeAndNumericData()
	csv.ConvertCsvBySpecifyingDelimiter()
	csv.ConvertCsvBySpecifyingEncoding()
	email.ConvertEmailWithAlteringFieldsVisibility()
	email.ConvertEmailWithAttachments()
	email.ConvertEmailWithFieldLabels()
	email.ConvertEmailWithTimezoneOffset()
	html.ConvertHtmlWithPageNumbering()
	note.ConvertNoteBySpecifyingFontSubstitution()
	pdf.ConvertPdfAndFlattenAllFields()
	pdf.ConvertPdfAndHideAnnotations()
	pdf.ConvertPdfAndRemoveEmbeddedFiles()
	presentation.ConvertPresentationByHidingComments()
	presentation.ConvertPresentationBySpecifyingFontSubstitution()
	presentation.ConvertPresentationWithHiddenSlidesIncluded()
	spreadsheet.ConvertSpreadsheetAndHideComments()
	spreadsheet.ConvertSpreadsheetByShowingGridLines()
	spreadsheet.ConvertSpreadsheetBySkippingEmptyRowsAndColumns()
	spreadsheet.ConvertSpreadsheetBySpecifyingFontsubstitution()
	spreadsheet.ConvertSpreadsheetBySpecifyingRange()
	spreadsheet.ConvertSpreadsheetWithHiddenSheetsIncluded()
	txt.ConvertTxtByControllingLeadingSpacesBehavior()
	txt.ConvertTxtByControllingTrailingSpacesBehavior()
	txt.ConvertTxtBySpecifyingEncoding()
	wordprocessing.ConvertWordProcessingByHidingComments()
	wordprocessing.ConvertWordProcessingByHidingTrackedChanges()
	wordprocessing.ConvertWordProcessingBySpecifyingFontSubstitution()
}
