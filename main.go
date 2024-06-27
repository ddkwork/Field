package main

import (
	"Field"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("Field", func(w *unison.Window) {
		w.Content().AddChild(field.New().Layout())
	})
}
