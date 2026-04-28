package main

import (
	"fmt"
	"time"
)

// Player represents the player character.
type Player struct {
	ID       int
	Name     string
	Physics  *physics.RigidBody
	Speed    float64
	JumpForce float64
}

func NewPlayer(id int, name string, initialX, initialY float64) *Player {
	return &Player{
		ID:       id,
		Name:     name,
		Physics:  physics.NewRigidBody(initialX, initialY),
		Speed:    5.0,       // Base horizontal speed
		JumpForce: 15.0,     // Initial jump force
	}
}

// Update handles player input and physics application.
func (p *Player) Update(deltaTime float64, inputX float64, inputJump bool) {
	// Horizontal movement (Momentum driven)
	moveVector := physics.Vector2{X: inputX * p.Speed, Y: 0}
	p.Physics.ApplyForce(moveVector)

	// Jumping
	if inputJump {
		p.Physics.ApplyForce(physics.Vector2{X: 0, Y: p.JumpForce})
	}

	// Apply physics updates
	p.Physics.Update(deltaTime)
}

func (p *Player) String() string {
	return fmt.Sprintf("Player {%s, Pos: (%.2f, %.2f), Vel: (%.2f, %.2f), Grounded: %t}",
		p.Name,
		p.Physics.Position.X,
		p.Physics.Position.Y,
		p.Physics.Velocity.X,
		p.Physics.Velocity.Y,
		p.Physics.IsGrounded)
}