package user

import "testing"

func TestUser(t *testing.T) {
	var testEqualString = func(expected, got string, t *testing.T) {
		t.Helper()
		if expected != got {
			t.Errorf("expected '%s' got '%s'\n", expected, got)
		}
	}

	t.Run("create a new User", func(t *testing.T) {
		user := User{
			ID:       1,
			Username: "Demo Article",
			Password: "1,2,3,4",
		}

		got := user.Username
		expected := "Demo Article"
		testEqualString(expected, got, t)

		got = user.Password
		expected = "1,2,3,4"
		testEqualString(expected, got, t)
	})
}
