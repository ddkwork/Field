package main

import (
	"Field"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("Field", func(w *unison.Window) {
		field.New().Layout(w.Content())
	})
}
