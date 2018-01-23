package observer

import "testing"

func TestSubject(t *testing.T) {
	wd := NewWeatherData()
	wd.setMeasurements(wd.temperature, 95.0, wd.pressure)
}

func TestObserver(t *testing.T) {

	wd := NewWeatherData()

	_ = NewCurrentConditionsDisplay("current", wd)
	_ = NewCurrentConditionsDisplay("hoge", wd)
	_ = NewCurrentConditionsDisplay("huga", wd)

	wd.setMeasurements(27, 65, 30)
	wd.setMeasurements(28, 75, 30)
	wd.setMeasurements(24, 35, 30)
}
