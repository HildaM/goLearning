package golang_2

import (
	"fmt"
	"strconv"
)

// 1. 定义接口
type Phone1 interface {
	call()
}

// 2. 实现接口
type Iphone struct {
	name string
}

func (iphoneSe Iphone) call() {
	fmt.Println("This is iphoneSE")
}

// 3. 使用接口实现多态

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// Phone实现
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// FreeGift
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

// 计算所有价格
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}

	return allPrice
}

func Interface1() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
