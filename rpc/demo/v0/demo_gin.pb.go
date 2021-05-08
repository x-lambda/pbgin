// Code generated by github.com/x-lambda/protoc-gen-gin. DO NOT EDIT.

package demo_v0

import (
	context "context"
	errors "errors"

	gin "github.com/gin-gonic/gin"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the github.com/x-lambda/protoc-gen-gin package it is being compiled against.
// context.
// gin.errors.

type BlogServiceHTTPServer interface {
	CreateArticle(context.Context, *Article) (*Article, error)

	GetArticles(context.Context, *GetArticlesReq) (*GetArticlesResp, error)
}

func RegisterBlogServiceHTTPServer(r gin.IRouter, srv BlogServiceHTTPServer) {
	s := BlogService{
		server: srv,
		router: r,
		resp:   defaultBlogServiceResp{},
	}
	s.RegisterService()
}

type BlogService struct {
	server BlogServiceHTTPServer
	router gin.IRouter
	resp   interface {
		Error(ctx *gin.Context, err error)
		ParamsError(ctx *gin.Context, err error)
		Success(ctx *gin.Context, data interface{})
	}
}

// Resp 返回值
type defaultBlogServiceResp struct{}

func (resp defaultBlogServiceResp) response(ctx *gin.Context, status, code int, msg string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Error 返回错误信息
func (resp defaultBlogServiceResp) Error(ctx *gin.Context, err error) {
	code := -1
	status := 500
	msg := err.Error()

	if err == nil {
		msg += ", err is nil"
		resp.response(ctx, status, code, msg, nil)
		return
	}

	type iCode interface {
		HTTPCode() int
		Message() string
		Code() int
	}

	var c iCode
	if errors.As(err, &c) {
		status = c.HTTPCode()
		code = c.Code()
		msg = c.Message()
	}

	_ = ctx.Error(err)

	resp.response(ctx, status, code, msg, nil)
}

// ParamsError 参数错误
func (resp defaultBlogServiceResp) ParamsError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	resp.response(ctx, 400, 400, "参数错误", nil)
}

// Success 返回成功信息
func (resp defaultBlogServiceResp) Success(ctx *gin.Context, data interface{}) {
	resp.response(ctx, 200, 0, "success", data)
}

func (s *BlogService) GetArticles_0(ctx *gin.Context) {
	var in GetArticlesReq

	if err := ctx.ShouldBindUri(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	if err := ctx.ShouldBindQuery(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	// TODO use metadata.NewIncomingContext
	// support GRPC/HTTP
	// timeout check with newCtx.Done()
	// md := metadata.New(nil)
	// for k, v := range ctx.Request.Header {
	// md.Set(k, v...)
	// }
	// newCtx := metadata.NewIncomingContext(ctx, md)
	newCtx := ctx.Request.Context()
	out, err := s.server.(BlogServiceHTTPServer).GetArticles(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) GetArticles_1(ctx *gin.Context) {
	var in GetArticlesReq

	if err := ctx.ShouldBindQuery(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	// TODO use metadata.NewIncomingContext
	// support GRPC/HTTP
	// timeout check with newCtx.Done()
	// md := metadata.New(nil)
	// for k, v := range ctx.Request.Header {
	// md.Set(k, v...)
	// }
	// newCtx := metadata.NewIncomingContext(ctx, md)
	newCtx := ctx.Request.Context()
	out, err := s.server.(BlogServiceHTTPServer).GetArticles(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) CreateArticle_0(ctx *gin.Context) {
	var in Article

	if err := ctx.ShouldBindUri(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&in); err != nil {
		s.resp.ParamsError(ctx, err)
		return
	}

	// TODO use metadata.NewIncomingContext
	// support GRPC/HTTP
	// timeout check with newCtx.Done()
	// md := metadata.New(nil)
	// for k, v := range ctx.Request.Header {
	// md.Set(k, v...)
	// }
	// newCtx := metadata.NewIncomingContext(ctx, md)
	newCtx := ctx.Request.Context()
	out, err := s.server.(BlogServiceHTTPServer).CreateArticle(newCtx, &in)
	if err != nil {
		s.resp.Error(ctx, err)
		return
	}

	s.resp.Success(ctx, out)
}

func (s *BlogService) RegisterService() {

	s.router.Handle("GET", "/v1/author/:author_id/articles", s.GetArticles_0)

	s.router.Handle("GET", "/v1/articles", s.GetArticles_1)

	s.router.Handle("POST", "/v1/author/:author_id/articles", s.CreateArticle_0)

}
