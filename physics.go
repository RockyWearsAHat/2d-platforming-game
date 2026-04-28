// /Users/alexwaldmann/Desktop/MyEditor/engine/projects/RockyWearsAHat-2d-platforming-game/physics.go
package main

import "math"

type Vector2 struct {
	X, Y float64
}

type RigidBody struct {
	Position Vector2
	Velocity Vector2
	Acceleration Vector2
	Mass float64
	IsGrounded bool
}

func NewRigidBody(x, y float64) *RigidBody {
	return &RigidBody{
		Position: Vector2{X: x, Y: y},
		Velocity: Vector2{X: 0, Y: 0},
		Acceleration: Vector2{X: 0, Y: 0},
		Mass: 1.0,
		IsGrounded: false,
	}
}

func (rb *RigidBody) ApplyForce(force Vector2) {
	// F = ma => a = F/m
	rb.Acceleration.X += force.X / rb.Mass
	rb.Acceleration.Y += force.Y / rb.Mass
}

func (rb *RigidBody) Update(dt float64) {
	// Apply acceleration to velocity
	rb.Velocity.X += rb.Acceleration.X * dt
	rb.Velocity.Y += rb.Acceleration.Y * dt

	// Apply velocity to position
	rb.Position.X += rb.Velocity.X * dt
	rb.Position.Y += rb.Velocity.Y * dt

	// Simple gravity simulation (assuming Y is up/down for simplicity in this core)
	// Let's assume Y increases downwards for screen coordinates, hence gravity pulls down.
	const gravity = 9.81 // Example gravity constant
	rb.Acceleration.Y += gravity / rb.Mass
}