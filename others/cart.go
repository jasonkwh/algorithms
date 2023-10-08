package others

var customerDiscount map[CustomerType]int = map[CustomerType]int{
	CustomerTypeNormal: 0,
	CustomerTypeSenior: 50,
}

type ShoppingCart struct {
	Items    []ShoppingItem
	Customer Customer
}

func (c ShoppingCart) GetTotalPrice() int {
	total := 0
	for _, item := range c.Items {
		total += item.GetItemPrice()
	}

	return calculateDiscountPrice(total, customerDiscount[c.Customer.Type])
}

type Customer struct {
	Name string
	Type CustomerType
}

type ShoppingItem struct {
	Name  string
	Cents int
	Off   int
}

func (i *ShoppingItem) UpdateItemPrice(price int) {
	i.Cents = price
}

func (i ShoppingItem) GetItemPrice() int {
	return i.Cents
}

func (i ShoppingItem) GetItemActualPrice() int {
	return calculateDiscountPrice(i.Cents, i.Off)
}

func calculateDiscountPrice(cents, off int) int {
	return cents / 100 * (100 - off)
}

type CustomerType uint

const (
	CustomerTypeNormal = iota
	CustomerTypeSenior
)
