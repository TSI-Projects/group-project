package apiErrors

const (
	INTERNAL_ERROR_CODE         = "INTERNAL_ERROR"
	FIELD_VALIDATION_ERROR_CODE = "FIELD_VALIDATION_ERROR"
	ID_NOT_SPECIFIED_ERROR_CODE = "ID_NOT_SPECIFIED_ERROR"
	NOT_FOUND_ERROR_CODE        = "NOT_FOUND_ERROR"
)

const (
	DEFAULT_ERROR_MESSAGE          = "We're sorry, something went wrong on our end. Please try again later."
	FIELD_VALIDATION_ERROR_MESSAGE = "Oops! It seems like you missed some fields. Please make sure all fields are filled out and try again."
	NOT_FOUND_ERROR_MESSAGE        = "The endpoint URL is invalid or missing. Please ensure you're using the correct API endpoint."
	ID_NOT_SPECIFIED_ERROR_MESSAGE = "ID is not specified."
	INTERNAL_ERROR_MESSAGE         = "Internal Error"
)
