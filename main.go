// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/rituals

import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/cmd"
)

var (
	version = "2.1.4" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
