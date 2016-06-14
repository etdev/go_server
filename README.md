# go_server

Extremely basic static file server in go

### Usage

To start the server

```bash
go run server.go [-p port_num] [-r root_directory]
```

e.g. to run on port `5252` with `my_files` as the root directory,

```bash
go run server.go -p 5252 -r my_files
```

Defaults are `8000` and `public`, respectively

### File Organization

```bash
public/
  index.html
  css/
    styles.css
  js/
    js_file.js
```

Files in subfolders should be served automatically.
