package observer

import "fmt"

// Subject tells value has been changed to observer
type Subject interface {
	registerObserver(o Observer)
	removeObserver(o Observer)
	notifyObservers()
}

type Observer interface {
	update(temp, humidity, pressure float64)
}

type Display interface {
	display()
}

type WeatherData struct {
	Subject
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (wd *WeatherData) registerObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) removeObserver(o Observer) {

	for i := 0; i < len(wd.observers); i++ {
		if wd.observers[i] == o {
			if i == 0 {
				wd.observers = wd.observers[1:]
				return
			}
			if i == len(wd.observers)-1 {
				wd.observers = wd.observers[0:i]
				return
			}
			wd.observers = append(wd.observers[0:i], wd.observers[i+1:]...)
		}
	}
}

func (wd *WeatherData) notifyObsevers() {
	for _, o := range wd.observers {
		o.update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.notifyObsevers()
}

func (wd *WeatherData) setMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers:   []Observer{},
		temperature: 0,
		humidity:    0,
		pressure:    0,
	}
}

type CurrentConditionsDisplay struct {
	name        string
	temperature float64
	humidity    float64
	weatherdata Subject
}

func (ccd *CurrentConditionsDisplay) update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.display()
}

func (ccd *CurrentConditionsDisplay) display() {
	fmt.Printf("hello from %s\n", ccd.name)
	fmt.Printf("気温 %f\n", ccd.temperature)
	fmt.Printf("湿度 %f%%\n", ccd.humidity)
	fmt.Println("-----------------------")
}

func NewCurrentConditionsDisplay(n string, s Subject) *CurrentConditionsDisplay {
	ccd := &CurrentConditionsDisplay{
		name:        n,
		weatherdata: s,
	}
	s.registerObserver(ccd)
	return ccd
}
