package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"shortenLink/cmd/sequence-rpc/internal/svc"
	"shortenLink/cmd/sequence-rpc/shortenLink/rpc/sequence"

	"github.com/zeromicro/go-zero/core/logx"
)

const sqlReplaceIntoStub = `REPLACE INTO sequence (stub) VALUES ('a')`

type NextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NextLogic {
	return &NextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NextLogic) Next(in *sequence.Empty) (*sequence.NextResp, error) {
	// 发号器逻辑
	stmt, err := l.svcCtx.Conn.Prepare(sqlReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.Field("err", err))
		return nil, status.Error(codes.Aborted, "conn.Prepare failed")
	}
	res, err := stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec failed", logx.Field("err", err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	lid, err := res.LastInsertId()
	if err != nil {
		logx.Errorw("res.LastInsertId failed", logx.Field("err", err))
		return nil, status.Error(codes.Internal, "LastInsertId failed.")
	}
	return &sequence.NextResp{
		Value: uint64(lid),
	}, nil
}
