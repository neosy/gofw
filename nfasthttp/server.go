package nfasthttp

import (
	"github.com/neosy/gofw/nbasic"
	"github.com/valyala/fasthttp"
)

type StandardResponse struct {
	Message string `json:"message,omitempty"`
}

func (data *StandardResponse) json() []byte {

	json, _ := nbasic.StructToJSON(data)

	return json
}

func ResponseSuccess(ctx *fasthttp.RequestCtx, statusCode int, data []byte) {
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.SetContentType("application/json")
	if len(data) > 0 {
		ctx.Write(data)
	}
}

func ResponseSuccessOK(ctx *fasthttp.RequestCtx, data []byte) {
	ResponseSuccess(ctx, fasthttp.StatusOK, data)
}

func ResponseFail(ctx *fasthttp.RequestCtx, statusCode int, data []byte) {
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.SetContentType("application/json")
	if len(data) > 0 {
		ctx.Write(data)
	}
}

func ResponseFailDefault(ctx *fasthttp.RequestCtx, statusCode int, msg string) {
	data := StandardResponse{
		Message: msg,
	}

	ResponseFail(ctx, statusCode, data.json())
}

func ResponseSuccessDefault(ctx *fasthttp.RequestCtx, statusCode int, msg string) {
	data := StandardResponse{
		Message: msg,
	}

	ResponseSuccess(ctx, statusCode, data.json())
}

func ResponseSuccessOKDefault(ctx *fasthttp.RequestCtx, msg string) {
	data := StandardResponse{
		Message: msg,
	}

	ResponseSuccessOK(ctx, data.json())
}
