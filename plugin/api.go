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
    name, pass, domain string
}

type server struct {
    url, admin, adminpass string
    cookie
}

func (u user, s server) loginServer(err) {

    // Post to the iRedAdmin installation with the username and password of an admin passed from the user structure.
    resp, err := resty.R().
      SetQueryParams(map[string]string{
          "username": u.name,
          "password": u.pass,
      }).
      SetHeader("Accept", "application/json").
      Post(s.url)
    
    // Was the ping responsive?
    if err != nil {
        return(err)
    } else {
        // Store the auth cookie in the server structure
        s.cookie := resp.Cookies()
    }   
}

func (u user, s server) createUser() {
    
}

func (u user, s server) destroyUser() {
    
}

func (u user, s server) renewUser() {
    
}
