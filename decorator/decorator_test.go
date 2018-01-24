package decorator

import "testing"

func TestDecorator(t *testing.T) {
	e := NewEspresso()
	if e.getCost() != 1.99 {
		t.Errorf("NewEspresso returns wrong price. expect=%f, actual=%f", 1.99, e.getCost())
	}

	dr := NewDarkRoast()
	mdr := NewMocha(dr)
	mmdr := NewMocha(mdr)
	wmmdr := NewWhip(mmdr)

	if wmmdr.getCost() != 1.49 {
		t.Errorf("customized DarkRoast returns wrong price. expect=%f, actual=%f", 1.49, wmmdr.getCost())
	}
}
