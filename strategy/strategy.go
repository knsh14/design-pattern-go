package strategy

import "fmt"

// FlyBehaviour defines methods to fly
type FlyBehaviour interface {
	Fly() string
}

// FlyWithWings is implemetation of FlyBehaviour
type FlyWithWings struct {
}

// Fly is Implementation of FlyWithWings
func (f *FlyWithWings) Fly() string {
	return fmt.Sprint("羽で飛ぶ")
}

// FlyNoWay is implemetation of FlyBehaviour
type FlyNoWay struct {
}

// Fly is Implementation of FlyWithWings
func (f *FlyNoWay) Fly() string {
	return fmt.Sprint("飛べない")
}

// QuackBehaviour is interface about quack
type QuackBehaviour interface {
	Quack() string
}

// Quack is normal quack
type Quack struct {
}

// Quack is implemetnation of normal quack
func (q *Quack) Quack() string {
	return fmt.Sprint("ガーガー")
}

// Squack is normal quack
type Squack struct {
}

// Quack is implemetnation of squack
func (q *Squack) Quack() string {
	return fmt.Sprint("キューキュー")
}

// MuteQuack is normal quack
type MuteQuack struct {
}

// Quack is implemetnation of mute quack
func (q *MuteQuack) Quack() string {
	return fmt.Sprint("")
}

// MullardDuck is normal duck
type MullardDuck struct {
	FlyBehaviour
	QuackBehaviour
}

// NewMullardDuck is generate mullard duck
func NewMullardDuck() *MullardDuck {
	f := &FlyWithWings{}
	q := &Quack{}
	m := &MullardDuck{f, q}
	return m
}
