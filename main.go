package main

import (
    "os"
    "log"
    // Vault Package Dependencies
    "github.com/hashicorp/vault/helper/pluginutil"
    "github.com/hashicorp/vault/logical/plugins"
    // Plugin
    iredadmin "github.com/promaethius/Vault-iRedAdmin-Pro/plugin"
)

func main() {
	apiClientMeta := &pluginutil.APIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := pluginutil.VaultPluginTLSProvider(tlsConfig)

	err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: iredadmin.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
