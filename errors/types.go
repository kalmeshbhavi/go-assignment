package errors

// Transient Errors
var (
	TemporaryUnavailable StringCode = "TemporaryUnavailable"
	Canceled             StringCode = "Canceled"
	Timeout              StringCode = "Timeout"
)

// User Errors
var (
	InvalidRequest   StringCode = "InvalidRequest"
	NotFound         StringCode = "NotFound"
	Unauthorized     StringCode = "Unauthorized"
	PermissionDenied StringCode = "PermissionDenied"
	MethodNotAllowed StringCode = "MethodNotAllowed"
)

// Internal Errors
var (
	Internal      StringCode = "Internal"
	Unknown       StringCode = "Unknown"
	AlreadyExists StringCode = "AlreadyExists"
)

// Service Specific Errors
var (
	InvalidFileSize StringCode = "InvalidFileSize"
	InvalidFileType StringCode = "InvalidFileType"
)

// S3 specific errors
var (
	S3Timeout        StringCode = "S3TimeOut"
	S3ObjectNotFound StringCode = "S3ObjectNotFound"
)

var validErrors = map[Code]bool{
	TemporaryUnavailable: true,
	Canceled:             true,
	Timeout:              true,
	InvalidRequest:       true,
	NotFound:             true,
	Unauthorized:         true,
	PermissionDenied:     true,
	MethodNotAllowed:     true,
	Internal:             true,
	InvalidFileSize:      true,
	InvalidFileType:      true,
	S3Timeout:            true,
	S3ObjectNotFound:     true,
}

func getTransientErrors() []Code {
	return []Code{
		TemporaryUnavailable,
		Canceled,
		Timeout,
	}
}
