package validate

import "golang.org/x/crypto/bcrypt"

// 加密密码
func HashAndSalt(pwdStr string) (pwdHash string, err error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	pwdHash = string(hash)
	return
}

// 校验密码
func ComparePW(hashPwd string, plainPwd string) bool {
	byteHash := []byte(hashPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	return err == nil
}
