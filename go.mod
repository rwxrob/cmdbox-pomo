module github.com/rwxrob/cmdbox-pomo

go 1.17

require (
	github.com/rwxrob/cmdbox v0.7.5
	github.com/rwxrob/conf-go v1.1.1
)

replace github.com/rwxrob/conf-go => ../conf-go

replace github.com/rwxrob/cmdbox => ../cmdbox

require gopkg.in/yaml.v2 v2.4.0 // indirect
