package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dmahlow/go-bytearena/agent"
	"github.com/dmahlow/go-bytearena/models"
)

type Agent struct {
	agent.EmptyAgent
}

func (*Agent) Tick(perception models.Perception, _ int) *models.Actions {
	actions := &models.Actions{}
	motion := explorationVector(3)
	items := perception.Vision

	avoidObstacles(items, *motion)
	actions.Steer(motion.vector)
	return actions
}

func (*Agent) Error(err error, _ int) {
	fmt.Printf("Error received: %s\n", err.Error())
}

func avoidObstacles(items []*models.Item, v Vec) {
	for _, item := range items {
		iv := Vec{item.NearEdge}

		if iv.Distance() < 10 {
			if iv.Angle() < 180 {
				v.Set(-10, 0)
			} else {
				v.Set(10, 0)
			}
		}
	}
}

func explorationVector(y float64) *Vec {
	rand.Seed(time.Now().Unix())

	var d float64

	if rand.Float64() < 0.5 {
		d = 1
	} else {
		d = -1
	}

	x := d * rand.Float64()

	return &Vec{models.NewVector2(x, y)}
}
