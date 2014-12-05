GoBuildAll
==========

build main packages when you build sub packages


as we know, `go build` command onlyubuilds package in current directory
and in liteide, when editing a file of sub package, press ctrl-b only rebuilds the sub package, until you press ctrl-b in the editor of any file of main package.

 `gobuildall`  will search for `project.ini` in current folder,and then in parent folder , recursively , until reach your project root (eg. `$GOPATH\src\myproject\projcet.ini`)

if  `gobuildall`  didn't find `project.ini`, it will work just as `go`

if `project.ini` was found,
first `gobuildall` will build the current package
if fails, `gobuildall` will exit
if success, it will then build all packages listed in the .ini file


the parameters is fully compatiable with `go` command
stderr and stdout of `go` will be redirected to `gobuildall`

so you can use `gobuildall` instead of `go` in  any cases

currently, `gobuildall` only handles the `build` flag.
if you run `gobuildall` other than `gobuildall build` , it will run `go` with the same parameters.


sample project.ini

				[client]
				c:\\workspace\\src\\project\\client
				#pay attention that in windows,folder splitter is not \
				[server]
				c:\\workspace\\src\\project\\server


in liteide, you can modify `liteide\share\liteide\litebuild\gosrc.xml` by yourself to add/modify button settings.
you can add a `buildall` button , and then set hotkeys in liteide.
