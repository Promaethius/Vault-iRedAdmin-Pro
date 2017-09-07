package main

import (
    "os"

    "github.com/hashicorp/vault/helper/pluginutil"
    "github.com/hashicorp/vault/plugins"
)

func main() {
    apiClientMeta := &pluginutil.APIClientMeta{}
    flags := apiClientMeta.FlagSet()
    flags.Parse(os.Args)

    plugins.Serve(New().(iRedAdmin-Pro), apiClientMeta.GetTLSConfig())
}
