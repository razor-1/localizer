package cldr_test

import (
	"github.com/theplant/cldr"
	_ "github.com/theplant/cldr/resources/locales/en"

	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		locale string
		text   string
		data   []interface{}
		want   string
	}{
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 1}},
			want:   "1 item in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 2}},
			want:   "2 items in Your Cart in Your Cart",
		},
		{
			locale: "en",
			text:   `{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items in Your Cart")}} in Your Cart; {{p "Count2" (one "{{.Count2}} item") (other "{{.Count2}} items in Your Cart")}} in Your Cart`,
			data:   []interface{}{map[string]int{"Count": 2, "Count2": 1}},
			want:   "2 items in Your Cart in Your Cart; 1 item in Your Cart",
		},
		{
			locale: "en",
			text:   `{{$1}} {{$2}} {{$1}}`,
			data:   []interface{}{"string1", "string2"},
			want:   "string1 string2 string1",
		},
	}
	for i := 0; i < 2; i++ {
		for _, c := range cases {
			got, err := cldr.Parse(c.locale, c.text, c.data...)
			if err != nil {
				t.Error(err)
			}
			if got != c.want {
				t.Errorf("got %q; want %q", got, c.want)
			}
		}
	}
}
