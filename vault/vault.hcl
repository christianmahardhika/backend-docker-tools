backend "mysql" {
 address = "mcba-mysql:3306"
 username = "foo"
 password = "bar"
}

listener "tcp" {
 address = "0.0.0.0:8200"
 tls_disable = 1
}

ui = true

disable_mlock = true