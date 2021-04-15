package system

import "errors"

var ErrorSizesDoesNotMatch = errors.New("could not load full file")

var ErrorCreateFile = errors.New("could not create file")

var ErrorWriteFile = errors.New("could not write to file")

var ErrorLoading = errors.New("unexpected error occurred while loading file")

var ErrorOpening = errors.New("error while opening file")

var ErrorSeeking = errors.New("error while seeking file")