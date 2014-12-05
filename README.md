GoBuildAll
==========

build main packages when you build sub packages


as we know, `go build` command builds package in current directory
and in liteide, when editing a file of sub package, press ctrl-b only rebuilds the sub package, until you press ctrl-b in the editor of any file of main package.

GoBuildAll will recognize `project.ini` in your project root (eg. `$GOPATH\src\myproject\projcet.ini`)

first it will build the current package
if fails, GoBuildALl will exit
if success, it will then build all packages listed in the .ini file  


the parameters is fully compatiable with `go` command.
stderr and stdout of `go` will be redirected to GoBuildAll 

so you can just run `gobuildall build` as usual

 

sample project.ini

				[client]
				c:\\workspace\\src\\project\\client
				
				[server]
				c:\\workspace\\src\\project\\server
				
				
in liteide, you can modify `liteide\share\liteide\litebuild\gosrc.xml` by yourself to add/modify button settings.