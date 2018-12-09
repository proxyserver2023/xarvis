package article

import "testing"

func TestArticle(t *testing.T) {
	var testEqualString = func(expected, got string, t *testing.T) {
		t.Helper()
		if expected != got {
			t.Errorf("expected '%s' got '%s'\n", expected, got)
		}
	}

	t.Run("create a new Article", func(t *testing.T) {
		article := Article{
			ID:   1,
			Name: "Demo Article",
		}

		got := article.Name
		expected := "Demo Article"
		testEqualString(expected, got, t)
	})
}
