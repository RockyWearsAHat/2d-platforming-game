// main.go (Finalizing core logic based on scaffolding)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- RockyWearsAHat 2D Platforming Game ---")

	// 1. Setup Level
	level := NewLevel(800.0, 600.0, 100.0)
	fmt.Printf("Level initialized: Width=%.1f, Height=%.1f, GroundY=%.1f\n", level.Width, level.Height, level.GroundY)

	// 2. Setup Player
	player := NewPlayer(100.0, level.GroundY-50.0) // Start player slightly above the ground
	physics := NewPhysics(player)

	fmt.Printf("Player initialized at: (%.1f, %.1f)\n", player.X, player.Y)

	// 3. Simulation Loop Test
	fmt.Println("\nStarting simulation loop for 2 seconds...")
	startTime := time.Now()
	duration := 2 * time.Second
	frameTime := time.Duration(16 * time.Millisecond) // ~60 FPS tick as a placeholder

	for time.Since(startTime) < duration {
		deltaTime := frameTime
		physics.Update(deltaTime.Seconds())

		// Output state
		fmt.Printf("Time Elapsed: %.2fs | Pos: (%.2f, %.2f) | Vel: (%.2f, %.2f) | Momentum: %.2f | Grounded: %t\n",
			time.Since(startTime).Seconds(),
			player.X, player.Y,
			player.VelocityX, player.VelocityY,
			player.Momentum,
			player.IsGrounded)

		time.Sleep(frameTime)
	}

	fmt.Println("\nSimulation finished.")
}