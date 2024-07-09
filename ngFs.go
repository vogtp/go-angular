package angular

import (
	"io/fs"
	"log/slog"
	"net/http"
	"path"
)

// FileSystem creates a filesystem that is capable of serving angular web apps
func FileSystem(fs fs.FS) http.FileSystem {

	return &angularFS{
		slog: slog.Default().WithGroup("angularFS"),
		root: http.FS(fs),
	}
}

type angularFS struct {
	slog *slog.Logger
	root http.FileSystem
}

func (a *angularFS) Open(name string) (http.File, error) {
	a.slog.Debug("angularFS open", "name", name)
	f, err := a.root.Open(name)
	if err == nil {
		return f, nil
	}
	_, file := path.Split(name)
	name = "/" + file
	a.slog.Debug("angularFS try serve", "path", name)
	f, err = a.root.Open(name)
	if err == nil {
		return f, nil
	}
	name = "/index.html"
	a.slog.Debug("angularFS serve default", "path", name)
	return a.root.Open(name)
}
