package main

import "time"

// --- Level Structure ---
type Level struct {
	Width     float64
	Height    float64
	GroundY   float64
	Platforms []Platform
}

// --- Player Structure ---
type Player struct {
	X, Y         float64
	VelocityX    float64
	VelocityY    float64
	Momentum     float64 // For inertia/momentum physics
	IsGrounded   bool
	MaxSpeed     float64
	JumpStrength float64
}

// --- Physics Simulation ---
type Physics struct {
	Player *Player
	Gravity float64
	DeltaTime float64
}

func NewPlayer(startX, startY float64) *Player {
	return &Player{
		X:         startX,
		Y:         startY,
		VelocityX: 0,
		VelocityY: 0,
		Momentum:  0,
		IsGrounded: false,
		MaxSpeed:  15.0, // Base running speed
		JumpStrength: 10.0,
	}
}

func NewLevel(width, height, groundY float64) *Level {
	return &Level{
		Width:     width,
		Height:    height,
		GroundY:   groundY,
		Platforms: []Platform{
			{X: 0, Y: groundY}, // Ground level
		},
	}
}

type Platform struct {
	X, Y, Width, Height float64
}

func NewPhysics(p *Player) *Physics {
	return &Physics{
		Player: p,
		Gravity: 9.81, // Placeholder gravity (in game units)
		DeltaTime: 0.016666666666666666, // ~60 FPS base tick
	}
}

func (p *Player) ApplyGravity(dt float64) {
	if !p.IsGrounded {
		p.VelocityY += Physics.Gravity * dt
	}
}

func (p *Player) UpdateMovement(dt float64) {
	// Apply horizontal movement based on momentum
	p.X += p.VelocityX * dt

	// Apply vertical movement
	p.Y += p.VelocityY * dt

	// Apply friction/momentum damping
	p.VelocityX *= 0.95
	p.VelocityY *= 0.95

	// Check ground collision (simplified)
	p.IsGrounded = false
	for _, platform := range LevelBase.Platforms {
		// Simplified 2D collision check for demonstration
		if p.Y + 1 < platform.Y && p.Y+1 > platform.Y-platform.Height && p.X > platform.X && p.X < platform.X+platform.Width {
			p.Y = platform.Y - 1 // Snap to ground
			p.VelocityY = 0
			p.IsGrounded = true
			break
		}
	}
}

// Mock global level reference for context in the physics model
var LevelBase = &Level{}

func (p *Physics) Update(dt float64) {
	p.Player.ApplyGravity(dt)
	p.Player.UpdateMovement(dt)
}