/*******
* @Author:qingmeng
* @Description:
* @File:focus
* @Date2022/2/20
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/service"
	"shop/tool"
	"strconv"
)

//添加关注
func addFocus(ctx *gin.Context) {
	token := ctx.PostForm("token")
	iGoodsId := ctx.PostForm("goods_id")
	iStoreId := ctx.PostForm("store_id")
	fs := service.FocusService{}
	mc, _ := service.ParseToken(token)
	focus := model.Focus{
		Uid:     mc.User.Uid,
		GoodsId: 0,
		StoreId: 0,
	}

	if iStoreId == "" && iGoodsId == "" {
		tool.RespErrorWithData(ctx, "未输出参数")
		return
	}

	//验证goodsId
	if iGoodsId != "" {
		goodsId, err := strconv.Atoi(iGoodsId)
		if err != nil {
			fmt.Println("goodsId to int err:", err)
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}
		if goodsId <= 0 {
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}

		//是否已关注
		flag, err := fs.IsExistGoodsFocus(mc.User.Uid, goodsId)
		if err != nil {
			fmt.Println("judge exist focus err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if flag {
			tool.RespErrorWithData(ctx, "已关注过该商品")
			return
		}
		focus.GoodsId = goodsId
	}

	//验证storeId
	if iStoreId != "" {
		storeId, err := strconv.Atoi(iGoodsId)
		if err != nil {
			fmt.Println("goodsId to int err:", err)
			tool.RespErrorWithData(ctx, "sort_id无效")
			return
		}
		if storeId <= 0 {
			tool.RespErrorWithData(ctx, "store_id无效")
			return
		}
		//是否已关注
		flag, err := fs.IsExistStoreFocus(mc.User.Uid, storeId)
		if err != nil {
			fmt.Println("judge exist focus err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if flag {
			tool.RespErrorWithData(ctx, "已关注过该店铺")
			return
		}
		focus.StoreId = storeId
	}
	err := fs.AddFocus(focus)
	if err != nil {
		fmt.Println("add focus err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "关注成功")
}

//查看关注
func viewFocus(ctx *gin.Context) {
	token := ctx.PostForm("token")
	fs := service.FocusService{}
	mc, _ := service.ParseToken(token)

	focus, err := fs.SelectFocusByUid(mc.User.Uid)
	if err != nil {
		fmt.Println("select focus by uid err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if focus == nil {
		tool.RespSuccessfulWithData(ctx, "暂无关注")
		return
	}
	tool.RespSuccessfulWithData(ctx, focus)
}

//取消关注
func deleteFocus(ctx *gin.Context) {
	token := ctx.PostForm("token")
	iGoodsId := ctx.PostForm("goods_id")
	iStoreId := ctx.PostForm("store_id")
	fs := service.FocusService{}
	mc, _ := service.ParseToken(token)
	focus := model.Focus{
		Uid:     mc.User.Uid,
		GoodsId: 0,
		StoreId: 0,
	}

	if iStoreId == "" && iGoodsId == "" {
		tool.RespErrorWithData(ctx, "未输出参数")
		return
	}

	//验证goodsId
	if iGoodsId != "" {
		goodsId, err := strconv.Atoi(iGoodsId)
		if err != nil {
			fmt.Println("goodsId to int err:", err)
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}
		if goodsId <= 0 {
			tool.RespErrorWithData(ctx, "goods_id无效")
			return
		}

		//是否已关注
		flag, err := fs.IsExistGoodsFocus(mc.User.Uid, goodsId)
		if err != nil {
			fmt.Println("judge exist focus err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if !flag {
			tool.RespErrorWithData(ctx, "未关注过该商品")
			return
		}
		focus.GoodsId = 0
	}

	//验证storeId
	if iStoreId != "" {
		storeId, err := strconv.Atoi(iGoodsId)
		if err != nil {
			fmt.Println("goodsId to int err:", err)
			tool.RespErrorWithData(ctx, "sort_id无效")
			return
		}
		if storeId <= 0 {
			tool.RespErrorWithData(ctx, "store_id无效")
			return
		}
		//是否已关注
		flag, err := fs.IsExistStoreFocus(mc.User.Uid, storeId)
		if err != nil {
			fmt.Println("judge exist focus err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if flag {
			tool.RespErrorWithData(ctx, "未关注过该店铺")
			return
		}
		focus.StoreId = 0
	}

	err := fs.DeleteFocusByUid(mc.User.Uid)
	if err != nil {
		fmt.Println("delete focus err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx, "取消关注成功")
}
