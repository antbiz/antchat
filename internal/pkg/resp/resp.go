package resp

import (
	"github.com/antbiz/antchat/internal/pkg/errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Resp 统一的 RESTful 接口响应处理

// OK 常规数据的成功处理
func OK(r *ghttp.Request, data ...interface{}) {
	if len(data) == 0 {
		r.Response.WriteJsonExit(g.Map{})
	}
	r.Response.WriteJsonExit(data[0])
}

// PageOK 列表数据的成功处理
func PageOK(r *ghttp.Request, total int, data interface{}) {
	r.Response.WriteJsonExit(g.Map{
		"total": total,
		"data":  data,
	})
}

// Error 错误处理
func Error(r *ghttp.Request, err *errors.Error) {
	// TODO: 优化：格式化记录详细的出错信息，加入链路跟踪
	if err.OrigErr != nil {
		g.Log().Async().Error(err.OrigErr)
	}
	r.Response.WriteHeader(err.Code)
	r.Response.WriteJsonExit(err)
}

func BadRequest(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.BadRequest(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.BadRequest(reason, message))
	}
}

func Unauthorized(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.Unauthorized(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.Unauthorized(reason, message))
	}
}

func Forbidden(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.Forbidden(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.Forbidden(reason, message))
	}
}

func NotFound(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.NotFound(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.NotFound(reason, message))
	}
}

func Conflict(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.Conflict(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.Conflict(reason, message))
	}
}

func InternalServer(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.InternalServer(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.InternalServer(reason, message))
	}
}

func ServiceUnavailable(r *ghttp.Request, reason, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.ServiceUnavailable(reason, message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.ServiceUnavailable(reason, message))
	}
}

func InvalidArgument(r *ghttp.Request, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.InvalidArgument(message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.InvalidArgument(message))
	}
}

func AlreadyExists(r *ghttp.Request, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.AlreadyExists(message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.AlreadyExists(message))
	}
}

func DatabaseError(r *ghttp.Request, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.DatabaseError(message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.DatabaseError(message))
	}
}

func UnknownError(r *ghttp.Request, message string, err ...error) {
	if len(err) > 0 {
		Error(r, errors.UnknownError(message).WithOrigErr(err[0]))
	} else {
		Error(r, errors.UnknownError(message))
	}
}
