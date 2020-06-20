package carbon

type Switcher struct {
	Location
	Style   *Style
	Content []Element
	Current Element
}

func (switcher *Switcher) Rasterize() {
	if switcher.Style != nil {
		switcher.Style.Rasterize(switcher.Size())
	}
	for _, element := range switcher.Content {
		element.Rasterize()
	}
}

func (switcher *Switcher) Handle(event Event, x, y float64) {
	switcher.Current.Handle(event, x, y)
}

func (switcher *Switcher) Draw(win *Window) {
	x, y := switcher.Center()
	switcher.Style.Draw(win, x, y)
	switcher.Current.Draw(nil)
}
