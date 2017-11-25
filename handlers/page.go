package handlers

type page struct {
	Meta          meta    `json:"meta,omitempty"`
	Content       content `json:"content,omitempty"`
	CopyrightYear string  `json:"copyrightyear,omitempty"`
}

type content struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type meta struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}
