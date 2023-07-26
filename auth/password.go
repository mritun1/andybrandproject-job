package auth

import "golang.org/x/crypto/bcrypt"

//------------------------------------------------------------------------------------------
// PASSWORD TO HASH - START
//------------------------------------------------------------------------------------------

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ------------------------------------------------------------------------------------------
// PASSWORD TO HASH - END
// ------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------
// VERIFY PASSWORD - START
// ------------------------------------------------------------------------------------------
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

//------------------------------------------------------------------------------------------
// VERIFY PASSWORD - END
//------------------------------------------------------------------------------------------
