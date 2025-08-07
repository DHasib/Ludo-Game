package game

// Manager orchestrates Ludo games. In production this would track
// players, board state, turns, and rules such as collisions and win
// conditions. For brevity only the skeleton is provided.
type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}
