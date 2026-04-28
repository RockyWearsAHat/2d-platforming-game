// /Users/alexwaldmann/Desktop/MyEditor/engine/projects/RockyWearsAHat-2d-platforming-game/level.go
package main

import "fmt"

type Platform struct {
	X, Y float64
	Width, Height float64
}

type Level struct {
	Platforms []Platform
	Score int
}

func NewLevel() *Level {
	return &Level{
		Platforms: []Platform{},
		Score: 0,
	}
}

// Setup a simple level geometry
func (l *Level) AddPlatform(x, y, w, h float64) {
	l.Platforms = append(l.Platforms, Platform{X: x, Y: y, Width: w, Height: h})
}

// Placeholder collision check (very basic AABB check)
func (l *Level) CheckCollision(p *Player) (Platform, bool, float64) {
	for _, platform := range l.Platforms {
		// AABB collision check
		if p.Body.Position.X >= platform.X && p.Body.Position.X <= platform.X+platform.Width &&
			p.Body.Position.Y >= platform.Y && p.Body.Position.Y <= platform.Y+platform.Height {
			// Collision detected (Need to resolve this properly in a full system)
			return platform, true, 0 // Return platform at collision point
		}
	}
	return Platform{}, false, 0
}

// Placeholder for Score logic
func (l *Level) AddScore(points int) {
	l.Score += points
	fmt.Printf("Score updated! Current Score: %d\n", l.Score)
}