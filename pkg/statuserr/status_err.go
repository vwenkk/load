package statuserr

import (
	"github.com/vwenkk/load/pkg/ent"
	"github.com/vwenkk/load/pkg/i18n"
	"github.com/vwenkk/load/pkg/msg/logmsg"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type DBError struct {
	cause  error
	detail interface{}
}

func Wrap(err error, detail interface{}) error {
	return DBError{
		cause:  err,
		detail: detail,
	}
}

func (e DBError) Error() string {
	logx.Error(e.cause, e.detail)
	return e.cause.Error()
}

func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case DBError:
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", e.detail))
			return http.StatusOK, NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", e.detail))
			return http.StatusOK, NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return http.StatusOK, NewInternalError(i18n.DatabaseError)
		}
	default:
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return http.StatusOK, NewInternalError(i18n.DatabaseError)
	}
}

func NewInternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

func NewInvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

func NewNotFoundError(msg string) error {
	return status.Error(codes.NotFound, msg)
}

func NewAlreadyExistsError(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

func NewUnauthenticatedError(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}
