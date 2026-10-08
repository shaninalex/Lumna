package email

import (
	"gitlab.com/shaninalex/lumna/app/core/errs"
)

var (
	EntryWasNotFoundError      = errs.NotFound("EML00", "email entry was not found")
	EntryWasNotUpdatedError    = errs.Platform("EML01", "email entry was not updated")
	ProcessorUnableToTickError = errs.Platform("EML02", "email processor unable to process tick entries batch")
)
