//Instal the package bcrypt
//go get golang.org/x/crypto/bcrypt

package security

import "golang.org/x/crypto/bcrypt"

// Hash receives string and returns hash
func Hash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

//VerifyPasswd compares password and hash

func VerifyPasswd(passwdHash, passwdstring string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(passwdstring))
}
