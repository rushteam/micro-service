package handler

import (
	"context"
	errs "errors"
	"strconv"

	"gitee.com/rushteam/micro-service/common/utils/snowflake"

	"github.com/micro/go-micro/errors"
	// "gitee.com/rushteam/micro-service/common/utils"
	"gitee.com/rushteam/micro-service/common/pb/order_srv"
	"gitee.com/rushteam/micro-service/service/order-srv/model"
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	// "go.uber.org/zap"
)

//OrderService ...
type OrderService struct {
	Service micro.Service
	// logger *zap.Logger
}

//创建订单参数验证
func validateCreateReq(req *order_srv.CreateReq) error {
	// 商品数量负数判断
	if req.GetOrder() == nil {
		return errs.New("订单信息不完整")
	}
	order := req.GetOrder()
	//商品数量限制
	if len(order.Items) < 1 {
		return errs.New("商品数量不能为空")
	}
	//限制最大数量防止拖垮查询
	if len(order.Items) > 99 {
		return errs.New("商品数量不能超过99个")
	}
	for _, sku := range order.Items {
		if sku.SkuId <= 0 {
			return errs.New("选择商品异常")
		}
		if sku.Price < 0 {
			return errs.New("选择商品价格异常")
		}
		if sku.Qty < 1 {
			return errs.New("选择商品数量异常")
		}
	}
	return nil
}

//创建订单计算
func calOrder(req *order_srv.CreateReq) error {
	order := req.Order
	var skuIds []int64
	for _, sku := range order.Items {
		skuIds = append(skuIds, sku.SkuId)
	}
	// var orderList []*order_srv.Order
	//拉取sku todo 这里应该改为 product_srv服务调用
	Model := model.Db()
	skuList, err := Model.GetSkuListBySkuIds(skuIds)
	if err != nil || len(skuList) < 1 {
		log.Log("[error] OrderService.Create ", err.Error())
		return errs.New("选择的商品已失效")
	}
	var skuMaps = make(map[int64]*model.SkuModel, len(skuList))
	for i := range skuList {
		skuMaps[skuList[i].SkuID] = &skuList[i]
	}

	var orderTotal int64
	var orderDiscount int64
	var orderFreight int64
	var orderPayment int64

	for _, sku := range order.Items {
		stdSku, ok := skuMaps[sku.SkuId]
		if !ok {
			return errs.New("选择的商品已失效")
		}
		//计算总价
		if stdSku.Price != sku.Price {
			return errs.New("选择的商品价格发生变化")
		}

		//sku总价
		total := stdSku.Price * sku.Qty
		//sku运费
		freight := stdSku.Weight * 0

		//逻辑安全判断
		if sku.Qty < 0 {
			return errs.New("商品数量异常")
		}
		if stdSku.Price < 0 {
			return errs.New("商品价值异常")
		}
		if total < 0 {
			return errs.New("商品总价值异常")
		}
		if freight < 0 {
			return errs.New("运费值异常")
		}
		//赋值运费
		orderFreight += freight
		//赋值总价
		orderTotal += total
	}
	orderPayment += orderTotal
	orderPayment += orderFreight
	orderPayment -= orderDiscount

	order.Total = orderTotal
	order.Discount = orderDiscount
	order.Payment = orderPayment

	return nil
}

func genOrderNo(workerID uint32) (string, error) {
	s, err := snowflake.NewSnowFlake(workerID)
	if err != nil {
		return "", err
	}
	id, err := s.Next()
	if err != nil {
		return "", err
	}
	no := strconv.FormatUint(id, 10)
	return no, nil
}

//Create ...
func (s *OrderService) Create(ctx context.Context, req *order_srv.CreateReq, rsp *order_srv.OrderRsp) error {
	log.Log("[access] OrderService.Create")
	var err error
	err = validateCreateReq(req)
	if err != nil {
		return errors.BadRequest("OrderService.Create", err.Error())
	}
	//计算金额
	err = calOrder(req)
	if err != nil {
		return errors.BadRequest("OrderService.Create", err.Error())
	}
	//生成订单号
	orderNo, err := genOrderNo(1)
	if err != nil {
		return errors.BadRequest("OrderService.Create", err.Error())
	}
	req.Order.OrderNo = orderNo

	//保存订单
	// Model := model.Db()
	// err := Model.CreateOrder()

	rsp.Order = req.Order
	return nil
}

//Budget ..
func (s *OrderService) Budget(ctx context.Context, req *order_srv.CreateReq, rsp *order_srv.OrderRsp) error {
	log.Log("[access] OrderService.Budget")
	var err error
	err = validateCreateReq(req)
	if err != nil {
		return errors.BadRequest("OrderService.Budget", err.Error())
	}
	err = calOrder(req)
	if err != nil {
		return errors.BadRequest("OrderService.Budget", err.Error())
	}
	rsp.Order = req.Order
	return nil
}

//Order ..
func (s *OrderService) Order(ctx context.Context, req *order_srv.QueryReq, rsp *order_srv.OrderRsp) error {
	return nil
}
