package app

import "context"

// App struct
type App struct {
	Ctx context.Context
}

type MenuItem struct {
	Type     string `json:"type"` // item/separator
	Title    string `json:"title"`
	Tooltip  string `json:"tooltip"`
	Event    string `json:"event"`
	Disabled bool   `json:"disabled"`
}

type StatsSnapshot struct {
	ModifierStats map[string]uint64 `json:"modifierStats"`
	KeyStats      map[string]uint64 `json:"keyStats"`
	ComboStats    map[string]uint64 `json:"comboStats"`

	MouseLeftClick  uint64  `json:"mouseLeftClick"`
	MouseRightClick uint64  `json:"mouseRightClick"`
	MouseMovePixels float64 `json:"mouseMovePixels"`
	MouseWheel      int64   `json:"mouseWheel"`
}
