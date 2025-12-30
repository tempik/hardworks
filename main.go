package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)

func extraOvertravel(thickness, protrusion, diameter float64) float64 {
	R := diameter / 2
	yTop := R - protrusion
	yBot := R - protrusion + thickness
	xTop := math.Sqrt(math.Max(R*R-yTop*yTop, 0))
	xBot := math.Sqrt(math.Max(R*R-yBot*yBot, 0))
	return xTop - xBot
}

func travelForBottomLength(Lb, thickness, protrusion, diameter float64) float64 {
	return Lb + extraOvertravel(thickness, protrusion, diameter)
}

func main() {
	thickness := flag.Float64("t", 0, "толщина заготовки (мм)")
	protrusion := flag.Float64("d", 0, "вылет диска (мм)")
	diameter := flag.Float64("D", 0, "диаметр диска (мм)")
	Lb := flag.Float64("Lb", 0, "желаемая длина реза снизу (мм)")
	flag.Parse()

	if *thickness <= 0 || *protrusion <= 0 || *diameter <= 0 || *Lb < 0 {
		log.Fatal("Укажите положительные значения: -t -d -D [-Lb]")
	}

	delta := extraOvertravel(*thickness, *protrusion, *diameter)
	travel := travelForBottomLength(*Lb, *thickness, *protrusion, *diameter)

	fmt.Printf("Дополнительная длина хода: %.2f мм\n", delta)
	fmt.Printf("Полный ход пилы: %.2f мм\n", travel)
}
