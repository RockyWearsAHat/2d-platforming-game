package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("RockyWearsAHat 2D Platformer Simulation Started")

	// --- Setup ---
	player := NewPlayer(100.0, 100.0)
	level := NewLevel()

	// Setup level geometry
	// Create a ground platform at Y=500
	level.AddPlatform(0, 500, 800, 100)

	// --- Simulation Loop ---
	dt := 1.0 / 60.0 // Target 60 FPS

	fmt.Println("Starting simulation loop for 5 seconds...")
	startTime := time.Now()
	
	for time.Since(startTime) < 5 * time.Second {
		// 1. Input/Force Simulation (Simulate a jump mid-simulation if possible)
		// In a real system, this is where input is read. We simulate a jump attempt here.
		if player.Body.IsGrounded {
			// Simulate input that attempts a jump occasionally
			if time.Since(startTime).Seconds() > 1.5 && time.Since(startTime).Duration() < 2*time.Second {
				player.Jump()
			}
		}
		
		// 2. Physics Update
		player.Update(dt)

		// 3. Collision (Basic level check against the floor)
		if player.Body.Position.Y >= 500 {
			player.Body.Position.Y = 500.0
			player.Body.Velocity.Y = 0
			player.Body.IsGrounded = true
		}

		// 4. Basic Logging (To observe state)
		if int(time.Since(startTime).Seconds()*60) % 10 == 0 { // Log every 10 frames (approx 0.16s)
			fmt.Printf("Time: %.2fs | Pos: (%.1f, %.1f) | VelY: %.1f | Grounded: %t\n", 
				time.Since(startTime).Seconds(), 
				player.Body.Position.X, 
				player.Body.Position.Y, 
				player.Body.Velocity.Y,
				player.Body.IsGrounded)
		}

		time.Sleep(time.Duration(dt*1000*float64(time.Second)))
	}

	// --- Final State ---
	fmt.Println("\nSimulation Finished.")
	fmt.Printf("Final Player Position: (%.2f, %.2f)\n", player.Body.Position.X, player.Body.Position.Y)
	fmt.Printf("Final Score: %d\n", level.Score)
}