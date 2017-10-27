package setting

// Table is the DB table for this model
var Table = "settings"

// Setting is a key value object that describes how page functionality should act
type Setting struct {
	ID    string `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
