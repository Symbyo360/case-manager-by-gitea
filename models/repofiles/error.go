package repofiles

import (
	"fmt"

	"code.gitea.io/gitea/modules/util"
)

type ErrFileMetaAlreadyExist struct {
	Sha string
}

// IsErrFileAlreadyExist checks if an error is a IsErrFileAlreadyExist.
func IsErrFileMetaAlreadyExist(err error) bool {
	_, ok := err.(ErrFileMetaAlreadyExist)
	return ok
}

func (err ErrFileMetaAlreadyExist) Error() string {
	return fmt.Sprintf("File already exists [SHA: %s]", err.Sha)
}

// Unwrap unwraps this error as a ErrExist error
func (err ErrFileMetaAlreadyExist) Unwrap() error {
	return util.ErrAlreadyExist
}

type ErrFileMetaNotExist struct {
	Sha string
}

// IsErrFileAlreadyExist checks if an error is a ErrUAlreadyExists.
func IsErrFileMetaNotExist(err error) bool {
	_, ok := err.(ErrFileMetaNotExist)
	return ok
}

func (err ErrFileMetaNotExist) Error() string {
	return fmt.Sprintf("File not exists [SHA: %s]", err.Sha)
}

// Unwrap unwraps this error as a ErrExist error
func (err ErrFileMetaNotExist) Unwrap() error {
	return util.ErrNotExist
}
