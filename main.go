package main

import (
	"Field"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("Field", func(w *unison.Window) {
		w.Content().AddChild(field.New().Layout())
	})
}
