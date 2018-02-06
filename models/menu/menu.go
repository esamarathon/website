package menu

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Title  string
	Link   string
	NewTab bool
}

func Default() Menu {
	return Menu{
		[]MenuItem{
			MenuItem{"Home", "/", false},
			MenuItem{"News", "/news", false},
			MenuItem{"Schedule", "/schedule", false},
			MenuItem{"Donate", "https://www.speedrun.com/esa2017/donate", true},
			MenuItem{"Tickets", "https://esamarathon.eventbrite.com/", true},
			MenuItem{"Forum", "https://www.speedrun.com/ESA_Winter_2018/forum", true},
		},
	}
}
