package progress

import "sync"

// ProgressStyles contains the styles of the progressbars. Custom styles can
// easily be appended and used.
var ProgressStyles = map[string][]string{
	"double":        []string{" ", "="},
	"double-":       []string{" ", "-", "="},
	"single":        []string{" ", "-"},
	"parallelogram": []string{"▱", "▰"},
	"spaced-blocks": []string{"▯", "▮"},
	"block":         []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"},
	"":              []string{" ", "="},
}

// progressStyleMtx locks the map so there are no concurrent map-accesses
var progressStyleMtx sync.RWMutex

// DefineStyle allows registering a custom progress-style to use. If the name
// is already defined it is overwritten. Directly defining it is possible but
// not recommended.
func DefineStyle(name string, characters []string) {
	progressStyleMtx.Lock()
	defer progressStyleMtx.Unlock()

	ProgressStyles[name] = characters
}
