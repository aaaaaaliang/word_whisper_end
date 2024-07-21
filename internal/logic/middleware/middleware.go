package middleware

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
	"strings"
	"word_whisper_end/internal/service"
	"word_whisper_end/utility"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) JWTMiddleware(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatus(http.StatusUnauthorized, "Authorization header is required")
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		r.Response.WriteStatus(http.StatusUnauthorized, "Authorization header must be Bearer token")
		return
	}
	tokenString := parts[1]

	token, claims, err := utility.ParseJWT(tokenString)
	if err != nil || token == nil || !token.Valid {
		r.Response.WriteStatus(http.StatusUnauthorized, "Invalid or expired token")
		return
	}
	// Extract user ID from JWT claims and add it to the context
	userId, ok := claims["userId"].(float64)
	if !ok {
		r.Response.WriteStatus(http.StatusUnauthorized, "Invalid token claims")
		return
	}
	// Convert the userId to an integer and set it in the request's context
	ctx := context.WithValue(r.Context(), "userId", int(userId))
	r.SetCtx(ctx)

	r.Middleware.Next()
}

func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	// Set response header to support UTF-8
	r.Response.Header().Set("Content-Type", "application/json; charset=utf-8")

	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
