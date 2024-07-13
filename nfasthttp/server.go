package nfasthttp

import (
	"github.com/neosy/gofw/nbasic"
	"github.com/valyala/fasthttp"
)

type StandardFailResponse struct {
	Success    bool                    `json:"success"`
	ErrorsData []StandardFailErrorData `json:"errorsData,omitempty"`
}

type StandardFailErrorData struct {
	Message string `json:"message,omitempty"`
}

func (data *StandardFailResponse) json() []byte {

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

func ResponseFailStandard(ctx *fasthttp.RequestCtx, statusCode int, msg string) {
	data := StandardFailResponse{
		Success:    false,
		ErrorsData: []StandardFailErrorData{{Message: msg}},
	}

	ResponseFail(ctx, statusCode, data.json())
}
