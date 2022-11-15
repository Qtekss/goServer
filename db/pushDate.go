package db

import (
	"fmt"
	"goServer/api"
)

type DbCoures struct {
	//CouresPush map[string]DbCoure
	CouresPush map[string]DbCoure
}

type DbCoure struct {
	Data     string  `json:"Data"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
}

func DataPush(ps *api.Сourse) *DbCoures {
	Set := NewSyncMap()
	var Part DbCoure
	Part.Data = ps.Date
	Courses := ps.Valute

	for index, element := range Courses {
		fmt.Println(index)
		Part.Data = ps.Date
		Part.CharCode = element.CharCode
		Part.Nominal = element.Nominal
		Part.Name = element.Name
		Part.Value = element.Value
		Set.CouresPush[index] = Part
	}

	return Set
}

func NewSyncMap() *DbCoures {
	return &DbCoures{CouresPush: make(map[string]DbCoure)}
}
