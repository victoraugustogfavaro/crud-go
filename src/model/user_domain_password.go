package model

import (
	"crypto/md5"
	"encoding/hex"
)

// função de criptografar a senha
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
	// trocar a senha passada pela nova criptografada
}
