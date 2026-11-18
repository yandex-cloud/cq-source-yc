package client

import (
	"context"
	"errors"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// expectedCodes are the gRPC codes that say the client cannot see the
// resource rather than that the sync broke: the service account has no access
// to the cloud or folder, the resource is gone (an image of a disk), the cloud
// does not use the service (backup providers of quotas) or the installation no
// longer serves the method. They are reasons to skip, with the message the
// grouping of the logs is done by.
var expectedCodes = map[codes.Code]string{
	codes.PermissionDenied: "resolver skipped: permission denied",
	codes.NotFound:         "resolver skipped: not found",
	codes.Unimplemented:    "resolver skipped: not implemented",
}

// ClassifyError is the error classifier of the scheduler. An expected error
// (expectedCodes) is suppressed: plugin-sdk neither counts it nor sends it as
// a sync error, and it is logged here as a warning with a message per code,
// since the text of the error is unique to every request (its request ids).
// Every other error, e.g. Unavailable or an error of ours, is raised.
func ClassifyError(ctx context.Context, err error, event schema.ErrorEvent) bool {
	code, ok := expectedCode(err)
	if !ok {
		return false
	}

	base := zerolog.Ctx(ctx)
	if c, ok := event.Client.(*Client); ok {
		base = &c.Logger
	}
	logger := base.Warn().
		Str("table", event.Table.Name).
		Str("client", event.Client.ID()).
		Str("phase", event.Phase.String()).
		Str("grpc_code", code.String())
	if event.Column != nil {
		logger = logger.Str("column", event.Column.Name)
	}
	logger.Err(err).Msg(expectedCodes[code])
	return true
}

// expectedCode returns the code of err if it is expected. A joined error, e.g.
// of the quotas of every service of a cloud, is expected only if all of its
// errors are, and then has the code of the first one: status.Code would look
// at the first one only.
func expectedCode(err error) (codes.Code, bool) {
	for e := err; e != nil; e = errors.Unwrap(e) {
		if joined, ok := e.(interface{ Unwrap() []error }); ok {
			errs := joined.Unwrap()
			if len(errs) == 0 {
				return codes.OK, false
			}
			for _, err := range errs[1:] {
				if _, ok := expectedCode(err); !ok {
					return codes.OK, false
				}
			}
			return expectedCode(errs[0])
		}
	}
	code := status.Code(err)
	_, ok := expectedCodes[code]
	return code, ok
}
