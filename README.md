# 2D Platforming Game

An original 2D action-platformer focused on fast movement, precise timing, and chaotic set-piece combat.

The game takes high-level inspiration from the sharp speed/style feel of Katana Zero and the explosive combat energy of Broforce, while remaining explicitly original in its world, characters, mechanics, level design, and narrative.

Core pillars:
- Momentum-driven traversal and close-quarters action.
- Destructible arenas and chain-reaction hazards.
- Stylish score-based encounters that reward creativity over repetition.

## Technical constraints

- **Platform: web browser.** The game runs in a modern desktop browser (Chrome/Firefox/Safari) with no native build, no installer, no game engine.
- **Stack: vanilla HTML5 + JavaScript (TypeScript optional).** Use the Canvas 2D API for rendering. No Phaser, no PixiJS, no Babylon, no React, no game framework — write the loop, input, physics, and rendering from scratch.
- **Build tooling: minimal.** A single `package.json` with `npm run build` (bundles to static files) and `npm test` (runs unit tests with a lightweight runner like `vitest` or `node --test`). A simple dev server is fine; no Docker, no CI/CD scaffolding.
- **Verification: tests + a runnable index.html.** Every plan step's acceptance command is something like `npm test` showing a specific spec passing, or `npm run build` exiting 0. The final step must produce a working `dist/index.html` that loads and starts the game when opened in a browser.

The goal is fast iteration: small dependency surface, instant feedback, runs anywhere a browser runs.

---

@engine
