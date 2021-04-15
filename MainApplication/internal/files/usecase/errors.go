package usecase

import "errors"

var ErrorSizesDoesNotMatch = errors.New("could not load full file")

var ErrorCreateFile = errors.New("could not create file")

var ErrorWriteFile = errors.New("could not write to file")

var ErrorLoading = errors.New("unexpected error occurred while loading file")

var ErrorCouldNotGenID = errors.New("could not generate file id")

var ErrorCouldNotSaveFileInfo = errors.New("could not save file info")
