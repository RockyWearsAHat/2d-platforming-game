// physics.go
package main

import "math"

const (
	Gravity = 9.8  // Simplified gravity for demonstration
	MaxSpeed = 10.0
	Friction = 0.9  // Velocity decay factor
)

// Physics handles the simulation of the game world.
type Physics struct {
	Player *Player
	G      float64
}

// NewPhysics initializes the physics system.
func NewPhysics(p *Player) *Physics {
	return &Physics{
		Player: p,
		G:      Gravity,
	}
}

// Update applies physics simulation to the player.
func (p *Physics) Update(deltaTime float64) {
	// Apply gravity
	p.Player.VelocityY += p.G * deltaTime

	// Apply friction/damping to horizontal velocity
	p.Player.VelocityX *= Friction

	// Apply movement based on velocity
	p.Player.X += p.Player.VelocityX * deltaTime
	p.Player.Y += p.Player.VelocityY * deltaTime

	// Clamp velocity to MaxSpeed
	if math.Abs(p.Player.VelocityX) > MaxSpeed {
		p.Player.VelocityX = math.Copysign(MaxSpeed, p.Player.VelocityX)
	}
	if math.Abs(p.Player.VelocityY) > MaxSpeed {
		p.Player.VelocityY = math.Copysign(MaxSpeed, p.Player.VelocityY)
	}

	// Basic ground detection placeholder (needs refinement)
	if p.Player.Y > 100.0 { // Placeholder for ground collision
		p.Player.Y = 100.0
		p.Player.VelocityY = 0.0
		p.Player.IsGrounded = true
	} else {
		p.Player.IsGrounded = false
	}

	// Momentum calculation (simplified)
	if p.Player.VelocityX != 0 || p.Player.VelocityY != 0 {
		p.Player.Momentum += math.Sqrt(p.Player.VelocityX*p.Player.VelocityX + p.Player.VelocityY*p.Player.VelocityY)
	}
}