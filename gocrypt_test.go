package gocrypto

import "testing"

func TestEncryptAndCheckPassword(t *testing.T) {
	encryptedPassword, _ := EncryptPassword("mYSecrEtpAssWorD!")

	result, _ := CheckPassword("mYSecrEtpAssWorD!", encryptedPassword)
	if result {
		t.Log("CheckPassword(\"mYSecrEtpAssWorD!\", ePassword) PASSED, returns true")
	} else {
		t.Error("CheckPassword(\"mYSecrEtpAssWorD!\", ePassword) FAILED, should return true, but returned false")
	}

	result, _ = CheckPassword("invalid password", encryptedPassword)
	if !result {
		t.Log("CheckPassword(\"invalid password\", ePassword) PASSED, returns false")
	} else {
		t.Error("CheckPassword(\"invalid password\", ePassword) FAILED, should return false, but returned true")
	}

	otherDifferentPassword, _ := EncryptPassword("OtherPassword")

	result, _ = CheckPassword("mYSecrEtpAssWorD!", otherDifferentPassword)
	if !result {
		t.Log("CheckPassword(\"mYSecrEtpAssWorD!\", otherDifferentPassword) PASSED, returns false")
	} else {
		t.Error("CheckPassword(\"mYSecrEtpAssWorD!\", otherDifferentPassword) FAILED, should return false, but returned true")
	}

}

func TestRandString(t *testing.T) {
	firstRandString := RandString(10)

	if len(firstRandString) == 10 {
		t.Log("Length of RandString(10) PASSED, is 10")
	} else {
		t.Errorf("Length of RandString(10) PASSED, should be 10, but returned %d", len(firstRandString))
	}
}
