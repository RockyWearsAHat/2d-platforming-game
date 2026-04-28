package main

import (
	"testing"
	"time"
)

func TestPlayerMotion(t *testing.T) {
	// Setup environment
	gameLevel := &Level{Width: 800, Height: 600}
	player := NewPlayer(1, "TestPlayer", 50.0, 50.0)
	
	// Mock physics update to force movement and gravity simulation
	deltaTime := 0.016666666666666667 // ~60 FPS
	
	// Initial state check
	initialPos := player.Physics.Position
	
	// Apply a large horizontal force (simulating a push)
	// We bypass the 'Update' which includes gravity/ground checks, and directly test ApplyForce accumulation based on the physics file logic.
	player.Physics.ApplyForce(physics.Vector2{X: 500.0, Y: 0}) // Apply a strong push

	// Update physics over time
	player.Physics.Update(deltaTime, 0, false)

	// Check if movement occurred (Velocity should have changed)
	if player.Physics.Velocity.X == 0.0 {
		t.Errorf("Expected horizontal velocity to change after applying force, got %.2f", player.Physics.Velocity.X)
	}
	
	// Check that Y position changed due to simulated gravity (should fall if not grounded, but our mock is simple)
	// Given the simple ground check in physics.go, Y should remain near 0 if it started at 50.0
	if player.Physics.Position.Y > initialPos.Y + 0.1 {
		t.Errorf("Y position unexpectedly increased: %.2f", player.Physics.Position.Y)
	}
	
	// Final check on grounded state
	if !player.Physics.IsGrounded {
		t.Logf("Player is not grounded after movement test. Pos: (%.2f, %.2f)", player.Physics.Position.X, player.Physics.Position.Y)
	}
}

func TestLevelInteraction(t *testing.T) {
	gameLevel := &Level{Width: 800, Height: 600}
	gameLevel.AddBlock("ground")
	
	// Test collision check (should pass for a valid position)
	if !gameLevel.CheckCollision(physics.Vector2{X: 100, Y: 100}) {
		t.Error("Collision check failed for a valid position.")
	}
}