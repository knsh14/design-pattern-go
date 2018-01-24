package decorator

// Beverage is base interface to calcurate
type Beverage interface {
	getDescription() string
	getCost() float64
}

// Drink is base of all drink
type Drink struct {
	description string
	cost        float64
}

func (d *Drink) getDescription() string {
	return d.description
}

func (d *Drink) getCost() float64 {
	return d.cost
}

// Espresso is bitter coffee
type Espresso struct {
	Drink
}

// NewEspresso returns new espresso
func NewEspresso() *Espresso {
	return &Espresso{Drink{
		description: "エスプレッソ",
		cost:        1.99,
	}}
}

// HouseBlend is normal coffee
type HouseBlend struct {
	Drink
}

// NewHouseBlend generates new house blend coffee
func NewHouseBlend() *HouseBlend {
	return &HouseBlend{Drink{
		description: "ハウスブレンドコーヒー",
		cost:        0.89,
	}}
}

// DarkRoast is normal coffee
type DarkRoast struct {
	Drink
}

// NewDarkRoast generates new house blend coffee
func NewDarkRoast() *DarkRoast {
	return &DarkRoast{Drink{
		description: "ダークローストコーヒー",
		cost:        0.99,
	}}
}

// CondimentDecorator is toppings
type CondimentDecorator struct {
	Drink
	beverage Beverage
}

func (c *CondimentDecorator) getDescription() string {
	return c.beverage.getDescription() + " +" + c.description
}

func (c *CondimentDecorator) getCost() float64 {
	return c.beverage.getCost() + c.cost
}

// Mocha is bitter cocoa
type Mocha struct {
	CondimentDecorator
}

// NewMocha generate new customize
func NewMocha(b Beverage) *Mocha {
	return &Mocha{CondimentDecorator{Drink: Drink{
		description: "モカ",
		cost:        0.2},
		beverage: b,
	}}
}

// Whip is Whipped Cream
type Whip struct {
	CondimentDecorator
}

// NewWhip generate new customize
func NewWhip(b Beverage) *Whip {
	return &Whip{CondimentDecorator{Drink: Drink{
		description: "ホイップ",
		cost:        0.1},
		beverage: b,
	}}
}
