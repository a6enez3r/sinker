# sync

sync local files

## quickstart

download the binary [for your platform] using curl

```sh
    curl -L  https://github.com/a6enez3r/sinker/raw/main/builds/sinker-darwin-amd64 >> sinker && chmod +x ./sinker
```

and run it

```sh
    ./sinker
```

you can see available commands using

```sh
usage: sinker [-h|--help] -c|--config "<value>" [-f|--filter "<value>"]

              sync files between local directories

Arguments:

  -h  --help    Print help information
  -c  --config  sync JSON configuration
  -f  --filter  filter to specify which items to sync. Default: all
```


## develop

```sh
cmds:

  help                  show usage / common commands available


              -- git --

  save-local:           save changes locally using git
  save-remote:          save changes to remote using git
  pull-remote:          pull changes from remote
  tag:                  create new tag, recreate if it exists

              -- go --

  deps:                 install deps [dev]
  build-single:         single platform build
  build-all:            cross platform build
  run:                  run package
  test:                 test package
  benchmark:            benchmark package
  coverage:             test coverage
  vet:                  vet modules

              -- code quality --

  lint:                 lint package
  format:               format package
  scan-duplicate:       scan package for duplicate code [dupl]
  scan-errors:          scan package for errors [errcheck]
  scan-security:        scan package for security issues [gosec]

              -- docker --

  build-env:            build docker env
  up-env:               start docker env
  exec-env:             exec. into docker env
  purge-env:            remove docker env
  status-env:           get status of docker env
  init-env:             init env + install common tools
```
