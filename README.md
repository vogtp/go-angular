# Golang Angular helpers

## angular.Filesystem

Server a static director (e.g. by embedding) as native go http server:

    //go:embed static
    var staticWeb embed.FS

    fsys, err := fs.Sub(staticWeb, "static")
    if err != nil {
    	panic(err)
    }
    http.Handle("/", http.FileServer(fsys))
    http.ListenAndServe(":8080", nil)
