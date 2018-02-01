package basic

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testcase := []struct {
		value, expect string
	}{
		{"Hello, world!", "!dlrow ,olleH"},
		{"Greetings from Earth", "htraE morf sgniteerG"},
		{"567329-2349", "9432-923765"},
		{"GoFreeCodeCamp", "pmaCedoCeerFoG"},
	}

	for _, c := range testcase {
		if got := ReverseString(c.value); got != c.expect {
			t.Errorf("Expected {{%s}}, got {{%s}}\n", c.expect, got)
		}
	}
}

// func TestReverseString(t *testing.T) {
// 	Convey("It should return the provided string in reversed order", t, func() {

// 		Convey("`Hello, world!` Should return => `htraE morf sgniteerG`", func() {
// 			test1 := ReverseString("Hello, world!")
// 			So(test1, ShouldEqual, "!dlrow ,olleH")
// 		})

// 		Convey("`Greetings from Earth` Should return => `!htraE morf sgniteerG`", func() {
// 			test2 := ReverseString("Greetings from Earth")
// 			So(test2, ShouldEqual, "htraE morf sgniteerG")
// 		})

// 		Convey("`567329-2349` Should return => `9432-923765`", func() {
// 			test3 := ReverseString("567329-2349")
// 			So(test3, ShouldEqual, "9432-923765")
// 		})

// 		Convey("`GoFreeCodeCamp` Should return =>  `pmaCedoCeerFoG`", func() {
// 			test4 := ReverseString("GoFreeCodeCamp")
// 			So(test4, ShouldEqual, "pmaCedoCeerFoG")
// 		})
// 	})
// }

// func TestFactorialize(t *testing.T) {
// 	Convey("It should return correct value of factorial", t, func() {

// 		Convey("Factorialize `5` Should return => `120`", func() {
// 			test1 := Factorialize(5)
// 			So(test1, ShouldEqual, 120)
// 		})

// 		Convey("Factorialize `10` Should return => `3628800`", func() {
// 			test2 := Factorialize(10)
// 			So(test2, ShouldEqual, 3628800)
// 		})

// 		Convey("Factorialize `20` Should return => `2432902008176640000`", func() {
// 			test3 := Factorialize(20)
// 			So(test3, ShouldEqual, 2432902008176640000)
// 		})

// 		Convey("Factorialize `0` Should return =>  `1`", func() {
// 			test4 := Factorialize(0)
// 			So(test4, ShouldEqual, 1)
// 		})
// 	})
// }

// func TestFactorializeWithErrors(t *testing.T) {
// 	Convey("It should return correct value of factorial and an error", t, func() {

// 		Convey("Factorialize `5` Should return => `120`", func() {
// 			test1, err := FactorializeWithError(5)
// 			So(test1, ShouldEqual, 120)
// 			So(err, ShouldBeNil)
// 		})

// 		Convey("Factorialize `10` Should return => `3628800`", func() {
// 			test2, err := FactorializeWithError(10)
// 			So(test2, ShouldEqual, 3628800)
// 			So(err, ShouldBeNil)
// 		})

// 		Convey("Factorialize `20` Should return => `2432902008176640000`", func() {
// 			test3, err := FactorializeWithError(20)
// 			So(test3, ShouldEqual, 2432902008176640000)
// 			So(err, ShouldBeNil)
// 		})

// 		Convey("Factorialize `0` Should return =>  `1`", func() {
// 			test4, err := FactorializeWithError(0)
// 			So(test4, ShouldEqual, 1)
// 			So(err, ShouldBeNil)
// 		})

// 		Convey("Factorialize `-6` Should return =>  `0` and throw error", func() {
// 			test5, err := FactorializeWithError(-6)
// 			So(test5, ShouldEqual, 0)
// 			So(err, ShouldNotBeNil)
// 		})
// 	})
// }

// func TestIsPalindrome(t *testing.T) {
// 	Convey(`It should remove all non-word charachters and then return true,
// 		if string is pallindrom`, t, func() {

// 		Convey("`_eye)` Should return => `true`", func() {
// 			test1 := IsPalindrome("_eye)")
// 			So(test1, ShouldBeTrue)
// 		})

// 		Convey("`race car` Should return => `true`", func() {
// 			test2 := IsPalindrome("race car")
// 			So(test2, ShouldBeTrue)
// 		})

// 		Convey("`A man, a plan, a canal. Panama` Should return => `true`", func() {
// 			test3 := IsPalindrome("A man, a plan, a canal. Panama")
// 			So(test3, ShouldBeTrue)
// 		})

// 		Convey("`nope` Should return =>  `false`", func() {
// 			test4 := IsPalindrome("nope")
// 			So(test4, ShouldBeFalse)
// 		})

// 		Convey("`not a palindrome` Should return =>  `false`", func() {
// 			test5 := IsPalindrome("not a palindrome")
// 			So(test5, ShouldBeFalse)
// 		})

// 		Convey("`1 eye for of 1 eye.` Should return =>  `false`", func() {
// 			test5 := IsPalindrome("1 eye for of 1 eye.")
// 			So(test5, ShouldBeFalse)
// 		})

// 		// If you are reading this, note that I'm using here backticks, so that go doesn't consider \
// 		// is an escape character
// 		Convey(`"0_0 (: /-\ :) 0-0" Should return =>  "true"`, func() {
// 			test6 := IsPalindrome(`0_0 (: /-\ :) 0-0`)
// 			So(test6, ShouldBeTrue)
// 		})

// 		Convey(`"five|\_/|four" Should return =>  "false"`, func() {
// 			test6 := IsPalindrome(`five|\_/|four`)
// 			So(test6, ShouldBeFalse)
// 		})
// 	})
// }

// func TestFindLongestWord(t *testing.T) {
// 	Convey("It should return correct value of factorial", t, func() {

// 		Convey("`May the force be with you` Should return => `5`", func() {
// 			test1 := FindLongestWord("May the force be with you")
// 			So(test1, ShouldEqual, 5)
// 		})

// 		Convey("`The quick brown fox jumped over the lazy dog` Should return => `6`", func() {
// 			test2 := FindLongestWord("The quick brown fox jumped over the lazy dog")
// 			So(test2, ShouldEqual, 6)
// 		})

// 		Convey("`Google do a barrel roll20` Should return => `6`", func() {
// 			test3 := FindLongestWord("Google do a barrel roll")
// 			So(test3, ShouldEqual, 6)
// 		})

// 		Convey("`What is the average airspeed velocity of an unladen swallow` Should return => `8`", func() {
// 			test4 := FindLongestWord("What is the average airspeed velocity of an unladen swallow")
// 			So(test4, ShouldEqual, 8)
// 		})

// 		Convey("`What if we try a super-long word such as otorhinolaryngology` Should return => `19`", func() {
// 			test4 := FindLongestWord("What if we try a super-long word such as otorhinolaryngology")
// 			So(test4, ShouldEqual, 19)
// 		})
// 	})
// }

// func TestTitleCase(t *testing.T) {
// 	Convey("It should return correct value of factorial", t, func() {

// 		Convey("`I'm a little tea pot` Should return => `I'm A Little Tea Pot`", func() {
// 			test1 := TitleCase("I'm a little tea pot")
// 			So(test1, ShouldEqual, "I'm A Little Tea Pot")
// 		})

// 		Convey("`sHoRt AnD sToUt` Should return => `Short And Stout`", func() {
// 			test2 := TitleCase("sHoRt AnD sToUt")
// 			So(test2, ShouldEqual, "Short And Stout")
// 		})

// 		Convey("`HERE IS MY HANDLE HERE IS MY SPOUT` Should return => `Here Is My Handle Here Is My Spout`", func() {
// 			test3 := TitleCase("HERE IS MY HANDLE HERE IS MY SPOUT")
// 			So(test3, ShouldEqual, "Here Is My Handle Here Is My Spout")
// 		})
// 	})
// }

// func TestConfirmEnding(t *testing.T) {
// 	Convey(`It should remove all non-word charachters and then return true,
// 		if string is pallindrom`, t, func() {

// 		Convey("`Bastian`, `n` Should return => `true`", func() {
// 			test1 := ConfirmEnding("Bastian", "n")
// 			So(test1, ShouldBeTrue)
// 		})

// 		Convey("`Connor`, `n` Should return => `false`", func() {
// 			test2 := ConfirmEnding("Connor", "n")
// 			So(test2, ShouldBeFalse)
// 		})

// 		Convey(`"Walking on water and developing software from
// 			a specification are easy if both are frozen", "specification" Should return => "true"`, func() {
// 			test3 := ConfirmEnding(`Walking on water and developing software from
// 				a specification are easy if both are frozen`, "specification")
// 			So(test3, ShouldBeFalse)
// 		})

// 		Convey("`He has to give me a new name`, `name` Should return =>  `true`", func() {
// 			test4 := ConfirmEnding("He has to give me a new name", "name")
// 			So(test4, ShouldBeTrue)
// 		})

// 		Convey("`Open sesame`, `same` Should return =>  `true`", func() {
// 			test5 := ConfirmEnding("Open sesame", "same")
// 			So(test5, ShouldBeTrue)
// 		})

// 		Convey("`Open sesame`, `pen` Should return =>  `false`", func() {
// 			test5 := ConfirmEnding("Open sesame", "pen")
// 			So(test5, ShouldBeFalse)
// 		})
// 	})
// }

// // 	Context("RepeatStringNumTimes", func() {
// // 		It("should repeat string n times without spaces", func() {
// // 			Expect(RepeatStringNumTimes("*", 3)).To(Equal("***"))
// // 			Expect(RepeatStringNumTimes("abc", 3)).To(Equal("abcabcabc"))
// // 			Expect(RepeatStringNumTimes("abc", 3)).To(Equal("abc"))
// // 			Expect(RepeatStringNumTimes("abc", -2)).To(Equal(""))
// // 		})
// // 	})

// // 	Context("TruncateString", func() {
// // 		It("should return string of n chars, that ends with ... if it was truncated", func() {
// // 			Expect(TruncateString("Peter Piper picked a peck of pickled peppers", 14)).To(Equal("Peter Piper..."))
// // 			Expect(TruncateString("A-", 1)).To(Equal("A..."))
// // 			Expect(TruncateString("Absolutely Longer", 2)).To(Equal("Ab..."))
// // 			Expect(TruncateString("abc", -2)).To(Equal("..."))
// // 			Expect(TruncateString("abc", 9)).To(Equal("abc"))
// // 		})
// // 	})

// // 	Context("ChunkArrayInGroups", func() {
// // 		ex1 := []int{1, 3, 5, 6, 2, 4, 7, 8}
// // 		out1 := [][]int{
// // 			{1, 3},
// // 			{5, 6},
// // 			{2, 4},
// // 			{7, 8},
// // 		}
// // 		ex2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// // 		out2 := [][]int{
// // 			{0, 1, 2, 3},
// // 			{4, 5, 6, 7},
// // 			{8},
// // 		}
// // 		It("should return string of n chars, that ends with ... if it was truncated", func() {
// // 			Expect(ChunkArrayInGroups(ex1, 2)).To(Equal(out1))
// // 			Expect(ChunkArrayInGroups(ex2, 4)).To(Equal(out2))
// // 		})
// // 	})
// // })
