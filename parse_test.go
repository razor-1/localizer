package cldr

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		locale string
		text   string
		data   interface{}
		want   string
	}{
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   map[string]int{"Count": 1},
			want:   "1 item in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   map[string]int{"Count": 2},
			want:   "2 items in Your Cart in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart; {{p "Count2" (one "{{.Count2}} item") (other "{{.Count2}} items in Your Cart")}} in Your Cart`,
			data:   map[string]int{"Count": 2, "Count2": 1},
			want:   "2 items in Your Cart in Your Cart; 1 item in Your Cart",
		},
	}
	for _, c := range cases {
		got, err := Parse(c.locale, c.text, c.data)
		if err != nil {
			t.Error(err)
		}
		if got != c.want {
			t.Errorf("got %q; want %q", got, c.want)
		}
	}
}
