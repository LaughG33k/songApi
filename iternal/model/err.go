package model

import "errors"

var (
	SongExists   = errors.New("song already exists")
	SongNotFound = errors.New("song not found")
	BadRequest   = errors.New("Bad request")
	IternalError = errors.New("Iternal error")
	EmptySong    = errors.New("empty song")
	EmptyGroup   = errors.New("empty group")
	EmptyText    = errors.New("empty text")
	EmptyLink    = errors.New("empty link")
	EmptyRealese = errors.New("empty realese date")
)
