package main

import (
    "os"
    // Vault Package Dependencies
    "github.com/hashicorp/vault/helper/pluginutil"
    "github.com/hashicorp/vault/plugins"
    // Go-Resty Dependency
    "gopkg.in/resty.v0"
)

// Define iRedAdmin as an interface to functions. Must satisfy user and server structs
type iRedAdmin interface {
    // One method to login and store the cookie for iRedAdmin-Pro
    loginServer()
    // Three methods to create, destroy, or renew a user within a domain.
    createUser()
    destroyUser()
    renewUser()
}

type user struct {
    username, password, domain string
}

type server struct {
    url, admin, adminpass string
}

func (u user, s server) loginServer() {
    
}

func (u user, s server) createUser() {
    
}

func (u user, s server) destroyUser() {
    
}

func (u user, s server) renewUser() {
    
}

func main() {
    apiClientMeta := &pluginutil.APIClientMeta{}
    flags := apiClientMeta.FlagSet()
    flags.Parse(os.Args)

    plugins.Serve(New().(iRedAdmin-Pro), apiClientMeta.GetTLSConfig())
}
