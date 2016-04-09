package schema

import "fmt"

type OpseeError interface {
	error

	// Returns the short phrase depicting the classification of the error.
	Code() string

	// Returns the error details message.
	Message() string

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErr() error
}

func (o OpseeBaseError) Code() string {
	return o.ErrorCode
}

func (o OpseeBaseError) Message() string {
	return o.ErrorMessage
}

func (o OpseeBaseError) Error() string {
	return fmt.Sprintf("ErrorCode: %s, ErrorMessage: %s", o.ErrorCode, o.ErrorMessage)
}

func (o OpseeBaseError) OrigErr() error {
	switch len(o.Errs) {
	case 0:
		return nil
	case 1:
		return o.Errs[0]
	default:
		errs := []error{}
		for _, err := range o.Errs {
			errs = append(errs, *err)
		}
		return NewOpseeBatchError("OpseeBatchedErrors", "multiple errors occured", errs)
	}
}

func (o OpseeBaseError) OrigErrs() []error {
	errs := []error{}
	for _, err := range o.Errs {
		errs = append(errs, *err)
	}
	return errs
}

func (o OpseeBaseError) OrigOpseeErr() *OpseeBaseError {
	if len(o.Errs) > 0 {
		return o.Errs[0]
	}
	return nil
}

func (o OpseeBaseError) OrigOpseeErrs() []*OpseeBaseError {
	return o.Errs
}

func NewOpseeBaseError(errorCode string, errorMessage string, opseeErrors []error) *OpseeBaseError {
	errs := []*OpseeBaseError{}
	for _, err := range opseeErrors {
		// test if err implements OpseeError interface (same as aws) for aws errors
		var i interface{} = err
		// otherwise no code, no base errors
		baseError := &OpseeBaseError{
			ErrorMessage: err.Error(),
		}

		if z, ok := i.(OpseeError); ok {
			baseError := &OpseeBaseError{
				ErrorCode:    z.Code(),
				ErrorMessage: z.Message(),
			}
			if r, ok := z.(OpseeBaseError); ok {
				baseError.Errs = r.OrigOpseeErrs()
			}

		}
		if z, ok := i.(OpseeBatchedError); ok {
			baseError = &OpseeBaseError{
				ErrorCode:    z.Code(),
				ErrorMessage: z.Message(),
			}
			if r, ok := z.(OpseeBaseError); ok {
				baseError.Errs = r.OrigOpseeErrs()
			}
		}
		errs = append(errs, baseError)
	}

	return &OpseeBaseError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		Errs:         errs,
	}
}

type OpseeBatchedError interface {
	OpseeError

	// Returns the original error if one was set.  Nil is returned if not set.
	OrigErrs() []error
}

func NewOpseeBatchError(code, message string, errs []error) OpseeBatchedError {
	return NewOpseeBaseError(code, message, errs)
}

func NewError(code, message string, origErr error) OpseeError {
	var errs []error
	if origErr != nil {
		errs = append(errs, origErr)
	}
	return NewOpseeBaseError(code, message, errs)
}
