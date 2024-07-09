# Golang Angular helpers

## angular.Filesystem

Server a static directory (e.g. by embedding) as native go http server:
import (
"embed"
"io/fs"
"net/http"

    "github.com/vogtp/go-angular"

)

    //go:embed static
    var staticWeb embed.FS

    fsys, err := fs.Sub(staticWeb, "static")
    if err != nil {
    	panic(err)
    }
    ngFS := angular.FileSystem(fsys)
    http.Handle("/", http.FileServer(ngFS))
    http.ListenAndServe(":8080", nil)
