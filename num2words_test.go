package num2words

import (
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("Small numbers should convert correctly", func(t *testing.T) {
		if have, want := Convert(0), "zero"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(1), "one"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(5), "five"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(10), "ten"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(11), "eleven"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(12), "twelve"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(17), "seventeen"; have != want {
			t.Fatal(have)
		}
	})
	t.Run("Tens should convert correctly", func(t *testing.T) {
		if have, want := Convert(20), "twenty"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(30), "thirty"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(40), "forty"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(50), "fifty"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(60), "sixty"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(90), "ninety"; have != want {
			t.Fatal(have)
		}
	})
	t.Run("Combined numbers should convert correctly", func(t *testing.T) {
		if have, want := Convert(21), "twenty-one"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(34), "thirty-four"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(49), "forty-nine"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(53), "fifty-three"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(68), "sixty-eight"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(99), "ninety-nine"; have != want {
			t.Fatal(have)
		}
	})
	t.Run("Big numbers should convert correctly", func(t *testing.T) {
		if have, want := Convert(100), "one hundred"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(200), "two hundred"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(500), "five hundred"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(123), "one hundred twenty-three"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(666), "six hundred sixty-six"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(1024), "one thousand twenty-four"; have != want {
			t.Fatal(have)
		}
	})
	t.Run("Negative numbers should convert correctly", func(t *testing.T) {
		if have, want := Convert(-123), "minus one hundred twenty-three"; have != want {
			t.Fatal(have)
		}
		if have, want := Convert(-99), "minus ninety-nine"; have != want {
			t.Fatal(have)
		}
	})
	t.Run("Convert with 'and' should convert correctly", func(t *testing.T) {
		if have, want := ConvertAnd(123), "one hundred and twenty-three"; have != want {
			t.Fatal(have)
		}
		if have, want := ConvertAnd(514), "five hundred and fourteen"; have != want {
			t.Fatal(have)
		}
		if have, want := ConvertAnd(1111), "one thousand one hundred and eleven"; have != want {
			t.Fatal(have)
		}
	})
}
