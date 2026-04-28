package physics

import "math"

const Gravity = 9.81 // Example gravity for simulation context

type Vector2 struct {
	X float64
	Y float64
}

type RigidBody struct {
	Position Vector2
	Velocity Vector2
	Mass     float64
	IsGrounded bool
}

func NewRigidBody(x, y float64) *RigidBody {
	return &RigidBody{
		Position: Vector2{X: x, Y: y},
		Velocity: Vector2{X: 0, Y: 0},
		Mass:     1.0,
		IsGrounded: false,
	}
}

// ApplyForce applies an impulse to the body.
func (rb *RigidBody) ApplyForce(force Vector2) {
	// Simplified: Force = Mass * Acceleration (a = F/m)
	acceleration := Vector2{X: force.X / rb.Mass, Y: force.Y / rb.Mass}
	rb.Velocity.X += acceleration.X
	rb.Velocity.Y += acceleration.Y
}

// Update applies gravity and updates position based on velocity.
func (rb *RigidBody) Update(deltaTime float64) {
	// Apply gravity
	gravityForce := Vector2{X: 0, Y: -Gravity} // Assuming Y is down in game terms or adjusting for positive Y up
	rb.ApplyForce(gravityForce)

	// Update position
	rb.Position.X += rb.Velocity.X * deltaTime
	rb.Position.Y += rb.Velocity.Y * deltaTime

	// Simple ground check placeholder (needs collision system)
	if rb.Position.Y > 0 { // Assuming ground is at Y=0
		rb.Position.Y = 0
		rb.Velocity.Y = 0
		rb.IsGrounded = true
	}
}