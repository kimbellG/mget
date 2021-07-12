package client

type processingString struct {
	percentInOneChar int
	buffer           []byte
}

func newProcessingString(percentInOneChar int) *processingString {
	return &processingString{
		percentInOneChar: percentInOneChar,
		buffer:           make([]byte, int(100/percentInOneChar)),
	}
}

func (ps *processingString) get(percentOfDownloaded int) string {

	for i := 0; i < len(ps.buffer); i++ {
		if ps.isSketchedChar(i, percentOfDownloaded) {
			ps.buffer[i] = '#'
		} else {
			ps.buffer[i] = ' '
		}
	}

	return string(ps.buffer)
}

func (ps *processingString) isSketchedChar(i, percentOfDownloaded int) bool {
	return (i+1)*ps.percentInOneChar < int(percentOfDownloaded)
}
