package strategy

import (
	"fmt"
	"testing"
)

func TestMullardDuck(t *testing.T) {
	var mullard interface{} = NewMullardDuck()

	f := &FlyWithWings{}
	if fb, ok := mullard.(FlyBehaviour); ok {
		if fb.Fly() != f.Fly() {
			t.Errorf("FlyBehaviour test failed. Expect = %s, Actual = %s", f.Fly(), fb.Fly())
		}
	}

	q := &Quack{}
	if qb, ok := mullard.(QuackBehaviour); ok {
		if qb.Quack() != q.Quack() {
			t.Errorf("QuackBehaviour test failed. Expect = %s. Actual = %s", q.Quack(), qb.Quack())
		}
	}

	duck := NewMullardDuck()
	fmt.Println(duck.Fly())
	fmt.Println(duck.Quack())
}
