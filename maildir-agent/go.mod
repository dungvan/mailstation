module github.com/dungvan/mailstation/maildir-agent

go 1.22.2

require (
	github.com/dungvan/mailstation/common v0.0.0-00010101000000-000000000000
	github.com/jhillyerd/enmime v1.2.0
	github.com/redis/go-redis/v9 v9.5.4
)

require (
	github.com/cention-sany/utf7 v0.0.0-20170124080048-26cad61bd60a // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gogs/chardet v0.0.0-20211120154057-b7413eaefb8f // indirect
	github.com/jaytaylor/html2text v0.0.0-20230321000545-74c2419ad056 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace github.com/dungvan/mailstation/common => ../common
