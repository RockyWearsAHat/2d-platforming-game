package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Mock update for the main execution flow
func main() {
	fmt.Println("--- Starting RockyWearsAHat-2D Platformer Simulation ---")

	// 1. Initialize Game World
	gameLevel := &Level{
		Width:  800.0,
		Height: 600.0,
		Blocks: make(map[string]bool),
	}
	gameLevel.AddBlock("ground")
	gameLevel.AddBlock("block_A")
	gameLevel.AddBlock("block_B")
	fmt.Printf("Level initialized: %s\n", gameLevel.Blocks)

	// 2. Initialize Player
	player := NewPlayer(1, "Rocky", 50.0, 50.0)
	fmt.Printf("Player initialized: %s\n", player.String())

	// 3. Simulation Loop (Mocking time progression)
	simDuration := 5 * time.Second
	currentTime := time.Now()
	lastTime := currentTime

	fmt.Println("\n--- Simulating Physics for 5 Seconds ---")

	for time.Since(lastTime) < simDuration {
		deltaTime := 1.0 / 60.0 // Simulate 60 FPS
		
		// Mock Player Input: Constant horizontal movement and occasional jump
		inputX := 0.0
		inputJump := false
		
		// Simulate movement: keep moving right and jump randomly
		if time.Since(currentTime) > 2*time.Second {
			inputX = 1.0 // Start moving right
		}
		if time.Since(currentTime) > 3*time.Second {
			inputJump = true // Try to jump
		}

		player.Update(deltaTime, inputX, inputJump)

		// Log state at the end of the frame
		if int(time.Since(currentTime).Seconds()*10) % 10 == 0 { // Log every ~10 frames
			fmt.Printf("Time: %.2fs | %s | Pos: (%.2f, %.2f) | Vel: (%.2f, %.2f) | Grounded: %t\n",
				time.Since(currentTime).Seconds(),
				player.Name,
				player.Physics.Position.X,
				player.Physics.Position.Y,
				player.Physics.Velocity.X,
				player.Physics.Velocity.Y,
				player.Physics.IsGrounded)
		}

		time.Sleep(time.Duration(deltaTime*float64(time.Second)))
		lastTime = time.Now()
	}

	fmt.Println("\n--- Simulation Complete ---")
}