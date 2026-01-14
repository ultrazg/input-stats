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
