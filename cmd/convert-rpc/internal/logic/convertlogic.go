package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shortenLink/cmd/sequence-rpc/shortenLink/rpc/sequence"
	"shortenLink/model"
	"shortenLink/pkg/base62"
	"shortenLink/pkg/md5"
	"shortenLink/pkg/urltool"

	"shortenLink/cmd/convert-rpc/internal/svc"
	"shortenLink/cmd/convert-rpc/shortenLink/rpc/convert"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConvertLogic) Convert(in *convert.ConvertRequest) (*convert.ConvertResponse, error) {
	longUrl := in.LongUrl
	//1. 判断长链接是否存在
	md5Value := md5.Sum([]byte(longUrl))
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, md5Value)
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, status.Error(codes.Aborted, "长链接已存在"+u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.Field("err", err), logx.Field("md5", md5Value))
		return nil, status.Error(codes.NotFound, "查询失败")
	}
	// 2. 输入的不能是短链接
	basePath, err := urltool.GetBasePath(longUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed,", logx.Field("err", err), logx.Field("longUrl", longUrl))
		return nil, status.Error(codes.Internal, err.Error())
	}

	surl, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, status.Error(codes.Aborted, "输入的不能是短链接")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.Field("err", err), logx.Field("surl", surl))
		return nil, status.Error(codes.NotFound, "查询失败")
	}
	// 3. 调用 sequence-rpc 获取新号
	var short string
	for {
		seqResp, err := l.svcCtx.SequenceRpc.Next(l.ctx, &sequence.Empty{})
		if err != nil {
			logx.Errorw("SequenceRpc.Next failed, ", logx.Field("err", err))
			return nil, status.Error(codes.Internal, err.Error())
		}
		seq := seqResp.Value
		// 3.1. 转换为短字符串
		//fmt.Printf("seqResp: %+v\n", seqResp)
		short = base62.Int2String(seq)
		// 3.2. 判断是否在黑名单中
		if _, ok := l.svcCtx.ShortUrlBlacklist[short]; !ok {
			break
		}
		logx.Infow("short in blacklist, get next sequence", logx.Field("short", short))
	}
	// 4. 存储映射关系
	if _, err := l.svcCtx.ShortUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		Md5:  md5Value,
		Lurl: sql.NullString{String: longUrl, Valid: true},
		Surl: sql.NullString{String: short, Valid: true},
	}); err != nil {
		logx.Errorw("ShortUrlModel.Insert failed", logx.Field("err", err), logx.Field("longUrl", longUrl), logx.Field("short", short))
		return nil, status.Error(codes.Internal, err.Error())
	}
	// 5. 更新布隆过滤器（忽略失败）
	if err := l.svcCtx.Filter.Add([]byte(short)); err != nil {
		logx.Errorw("bloom filter add failed", logx.Field("err", err), logx.Field("short", short))
	}
	// 6. 拼接完整短链并返回
	completeShortUrl := fmt.Sprintf("%s/%s", l.svcCtx.Config.ShortDomain, short)
	return &convert.ConvertResponse{
		ShortUrl: completeShortUrl,
	}, nil
}
