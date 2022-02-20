/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date2022/2/21
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/service"
	"shop/tool"
	"strconv"
	"time"
)

//添加评论
func addComment(ctx *gin.Context) {
	token := ctx.PostForm("token")
	iGoodsId := ctx.PostForm("goods_id")
	text := ctx.PostForm("text")
	iStar := ctx.PostForm("star")
	cs := service.CommentService{}
	gs := service.GoodsService{}
	ss := service.ShoppingCartService{}

	//检验参数
	if iGoodsId == "" || text == "" || iStar == "" {
		tool.RespErrorWithData(ctx, "参数无效")
		return
	}

	goodsId, err := strconv.Atoi(iGoodsId)
	if err != nil {
		fmt.Println("goods id to int err:", err)
		tool.RespErrorWithData(ctx, "goods_id无效")
		return
	}
	flag, err := gs.IsExistGoodsId(goodsId)
	if err != nil {
		fmt.Println("judge exist goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "goods_id无效")
		return
	}
	star, err := strconv.Atoi(iStar)
	if err != nil {
		fmt.Println("star id to int err:", err)
		tool.RespErrorWithData(ctx, "star无效")
		return
	}
	if star <= 0 || star > 5 {
		tool.RespErrorWithData(ctx, "star无效")
		return
	}

	flag = tool.CheckIfSensitive(text)
	if flag {
		tool.RespErrorWithData(ctx, "text无效")
		return
	}

	mc, _ := service.ParseToken(token)
	//检验用户是否对该商品已确认收货
	flag, err = ss.IsExistOrderByUidGoodsId(mc.User.Uid, goodsId)
	if err != nil {
		fmt.Println("judge exist order by uid and goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithData(ctx, "该用户对该商品未确认收货")
		return
	}

	//添加评论
	comment := model.Comment{
		GoodsId: goodsId,
		Uid:     mc.User.Uid,
		Text:    text,
		Star:    star,
		Date:    time.Now(),
	}
	err = cs.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//添加评论数
	err = gs.AddCommentNumber(goodsId)
	if err != nil {
		fmt.Println("add comment number err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//计算好评率
	goods, err := gs.SelectGoodsByGoodsId(goodsId)
	if err != nil {
		fmt.Println("select goods by goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}

	comments, err := cs.SelectCommentsByGoodsId(goodsId)
	if err != nil {
		fmt.Println("select comments by goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	stars := 0
	for _, c := range comments {
		stars += c.Star
	}
	feedback := float32(stars / goods.CommentNumber)

	err = gs.UpdateFeedback(goodsId, feedback)
	if err != nil {
		fmt.Println("update feedback err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "感谢您的评价")
}

//查看评论
func viewComment(ctx *gin.Context) {
	token := ctx.PostForm("token")
	goodsId := ctx.PostForm("goods_id")
	cs := service.CommentService{}
	var comments []model.Comment
	var gid int
	var err error

	if goodsId != "" {
		gid, err = strconv.Atoi(goodsId)
		if err != nil {
			fmt.Println("goodsId to int err:", err)
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}
		flag, err := cs.IsExistGoodsId(gid)
		if err != nil {
			fmt.Println("judge exist goods id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if !flag {
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}
	}

	if token == "" && goodsId == "" {
		tool.RespErrorWithData(ctx, "参数无效")
		return
	}

	if token != "" && goodsId == "" {
		mc, err := service.ParseToken(token)
		if err != nil {
			tool.RespErrorWithData(ctx, "token无效")
			return
		}
		comments, err = cs.SelectCommentsByUid(mc.User.Uid)
		if err != nil {
			fmt.Println("select comment by uid err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	if token != "" && goodsId != "" {
		mc, err := service.ParseToken(token)
		if err != nil {
			tool.RespErrorWithData(ctx, "token无效")
			return
		}
		comments, err = cs.SelectCommentsByUidGoodsId(mc.User.Uid, gid)
		if err != nil {
			fmt.Println("select comment by uid and goods id err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	if token == "" && goodsId != "" {
		comments, err = cs.SelectCommentsByGoodsId(gid)
		if err != nil {
			fmt.Println("select comment by goods id err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	if comments == nil {
		tool.RespSuccessfulWithData(ctx, "暂无评论")
		return
	}

	tool.RespSuccessfulWithData(ctx, comments)
}
