module github.com/node-real/nr-test-core

go 1.16

require (
	github.com/oliveagle/jsonpath v0.0.0-20180606110733-2e52cf6e6852
	github.com/stretchr/testify v1.7.2
)

replace github.com/stretchr/testify v1.7.2 => github.com/robertw07/testify v0.0.5

require (
	github.com/ethereum/go-ethereum v1.10.25
	github.com/google/go-cmp v0.5.5
	github.com/gorilla/websocket v1.5.0
	github.com/spf13/cobra v0.0.3
	github.com/tidwall/gjson v1.14.3
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/yaml.v3 v3.0.1
)
