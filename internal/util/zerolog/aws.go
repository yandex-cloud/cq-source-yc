// MIT License
//
// Copyright (c) 2022 Tim Voronov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// Original source https://github.com/ziflex/aws-zero/blob/5eac4aeb63e1879b63671a4b1ddb1c0483b877b5/logger.go

package zerolog

import (
	"github.com/aws/smithy-go/logging"
	"github.com/rs/zerolog"
)

type AWSLogger struct {
	log zerolog.Logger
}

// NewAWSLogger creates a Zerolog-based implementation of Smithy Logger.
func NewAWSLogger(log zerolog.Logger) logging.Logger {
	return &AWSLogger{log}
}

// Logf is expected to support the standard fmt package "verbs".
func (l *AWSLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	var evt *zerolog.Event

	switch classification {
	case logging.Warn:
		evt = l.log.Warn()
	case logging.Debug:
		evt = l.log.Debug()
	default:
		evt = l.log.Trace()
	}

	evt.Msgf(format, v...)
}
