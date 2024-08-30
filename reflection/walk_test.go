package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Mike"},
			[]string{"Mike"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mike", "Philadelphia"},
			[]string{"Mike", "Philadelphia"},
		},
		{
			"test with non string field",
			struct {
				Name string
				Age  int
			}{"Mike", 38},
			[]string{"Mike"},
		},
		{
			"nested fields",
			Person{
				"Mike",
				Profile{38, "Philadelphia"},
			},
			[]string{"Mike", "Philadelphia"},
		},
		{
			"pointers to things",
			&Person{
				"Mike",
				Profile{38, "{Philadelphia"},
			},
			[]string{"Mike", "Philadelphia"},
		},
		{
			"slices",
			[]Profile{
				{38, "Philadelphia"},
				{25, "South Africa"},
			},
			[]string{"Philadelphia", "South Africa"},
		},
		{
			"arrays",
			[2]Profile{
				{38, "Philadelphia"},
				{25, "South Africa"},
			},
			[]string{"Philadelphia", "South Africa"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got: %v Want: %v", got, test.ExpectedCalls)
			}
		})
	}
}
