package main

import (
    "os"
    "log"
    // Vault Package Dependencies
    "github.com/hashicorp/vault/helper/pluginutil"
    "github.com/hashicorp/vault/logical/plugins"
    // Go-Resty Dependency
    "gopkg.in/resty.v0"
)

func main() {
    apiClientMeta := &pluginutil.APIClientMeta{}
    flags := apiClientMeta.FlagSet()
    flags.Parse(os.Args)

    plugins.Serve(New().(iRedAdmin-Pro), apiClientMeta.GetTLSConfig())
}
