// /Users/alexwaldmann/Desktop/MyEditor/engine/projects/RockyWearsAHat-2d-platforming-game/player.go
package main

import "math"

const (
	MoveSpeed      = 300.0
	JumpStrength   = 500.0
	Gravity        = 9.81
	Friction       = 0.9
	AirControl     = 0.1  // Air resistance/control factor
)

type Player struct {
	Body *RigidBody
	JumpCount int
}

func NewPlayer(x, y float64) *Player {
	body := NewRigidBody(x, y)
	return &Player{
		Body: body,
		JumpCount: 0,
	}
}

func (p *Player) ApplyMovement(dx, dy float64) {
	// Horizontal movement
	acceleration := Vector2{X: dx / p.Body.Mass, Y: 0}
	p.Body.ApplyForce(acceleration)

	// Apply friction/damping for horizontal movement
	p.Body.Velocity.X *= Friction

	// Vertical movement (gravity handled in Update)
	p.Body.Velocity.Y += p.Body.Acceleration.Y // Apply current acceleration derived from gravity

	// Apply movement to position
	p.Body.Position.X += p.Body.Velocity.X * 0.016 // dt approximation (assuming 60 FPS -> dt approx 1/60 ~ 0.0167)
	p.Body.Position.Y += p.Body.Velocity.Y * 0.016
}

func (p *Player) Jump() bool {
	if p.Body.IsGrounded {
		p.Body.Velocity.Y = JumpStrength
		p.Body.IsGrounded = false
		p.JumpCount++
		return true
	}
	return false
}

func (p *Player) Update(dt float64) {
	// Apply gravity (only if not grounded)
	if !p.Body.IsGrounded {
		p.Body.Acceleration.Y += Gravity / p.Body.Mass
	}

	// Apply air control
	if !p.Body.IsGrounded {
		horizontalInput := 0.0
		if p.Body.Velocity.X != 0 {
			// Simulate horizontal control based on input (simplified: direct adjustment)
			// In a real game, this would involve reading input state.
			horizontalInput = math.Sin(p.Body.Velocity.X/100.0) * AirControl
			p.Body.Acceleration.X += horizontalInput / p.Body.Mass
		}
	}


	// Update position based on velocity
	p.Body.Position.X += p.Body.Velocity.X * dt
	p.Body.Position.Y += p.Body.Velocity.Y * dt

	// Simple ground check (needs proper collision handling later)
	if p.Body.Position.Y > 500 { // Arbitrary ground level check
		p.Body.Position.Y = 500
		p.Body.Velocity.Y = 0
		p.Body.IsGrounded = true
	}
}