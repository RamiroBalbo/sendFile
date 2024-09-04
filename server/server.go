package main

// Estructura para autenticar usuarios
type SimpleAuth struct{}

func (a *SimpleAuth) CheckPasswd(user, pass string) (bool, error) {
	if user == "usuario" && pass == "contraseña" {
		return true, nil
	}
	return false, nil
}

func main() {

}
