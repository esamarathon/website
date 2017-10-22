package handlers

type page struct {
	Meta    *meta
	Content *content
}

type content struct {
	Title string
	Body  string
}

type meta struct {
	Title       string
	Description string
	Image       string
}
