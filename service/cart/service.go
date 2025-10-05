package cart

import (
	"fmt"

	"github.com/Albert-tru/ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	ids := make([]int, 0, len(items))
	for _, item := range items {
		if item.ProductID <= 0 {
			return nil, fmt.Errorf("invalid product ID: %d", item.ProductID)
		}
		ids = append(ids, item.ProductID)
	}
	return ids, nil
}

// CreateOrder 创建订单，返回订单ID和总金额
func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	// 将产品列表转换为map，方便后续查找
	productMap := make(map[int]types.Product)
	for _, p := range ps {
		productMap[p.ID] = p
	}

	// 检查库存
	if err := checkStock(productMap, items); err != nil {
		return 0, 0, err
	}

	// 计算总价
	totalPrice := calculateTotalPrice(productMap, items)

	// 创建订单
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}

	// 创建订单项
	for _, cartItem := range items {
		p := productMap[cartItem.ProductID]
		err := h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: p.ID,
			Quantity:  cartItem.Quantity,
			Price:     p.Price,
		})
		if err != nil {
			return 0, 0, err
		}
	}

	return orderID, totalPrice, nil
}

func checkStock(productMap map[int]types.Product, items []types.CartItem) error {
	if len(productMap) == 0 {
		return fmt.Errorf("product map is empty")
	}
	for _, item := range items {
		p, exists := productMap[item.ProductID]
		if !exists {
			return fmt.Errorf("product ID %d not found", item.ProductID)
		}
		if p.Quantity < item.Quantity {
			return fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
		}
	}

	return nil
}

func calculateTotalPrice(productMap map[int]types.Product, items []types.CartItem) float64 {
	total := 0.0
	for _, item := range items {
		p, exists := productMap[item.ProductID]
		if !exists {
			continue
		}
		total += p.Price * float64(item.Quantity)
	}
	return total
}
