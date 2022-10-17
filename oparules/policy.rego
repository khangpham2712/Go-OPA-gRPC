package oparules

import future.keywords.if
import future.keywords.in

# By default, deny requests.
default allow := false

# Allow if that service doesn't require any grants.
# allow if { not (input.service in data.services) }

# Take jwt token and parse it to obtain the role of a user.
payload := p if {
	v := input.token
    io.jwt.verify_hs256(v, "dummy")
    [_, payload, _] := io.jwt.decode(v)
    p := payload
}

role := r if {
	r := payload.Role
}

# Allow if the user's grant equals to the required grant.
allow if {
	role == data.details[input.service][_]
}