package controller

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"github.com/cidermead/bjcp-api/include"
	"github.com/cidermead/bjcp-api/model"
)

// Style struct alias
type Style = model.Style

func insert(s []Style, at int, val Style) []Style {
	// Make sure there is enough room
	var v Style
	s = append(s, v)
	// Move all elements of s up one slot
	copy(s[at+1:], s[at:])
	// Insert the new element at the now free position
	s[at] = val
	return s
}

func Shuffle(a []Style) []Style {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}

func TrimCharacter(s string, c uint8) string {
	s1 := s
	if last := len(s1) - 1; last >= 0 && s1[last] == c {
		s1 = s1[:last]
	}
	return s1
}

func randBool() bool {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float32() < 0.5
	return r
}

func getRand(l int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(l)
	return r
}

func ShuffleAndCut(array []Style, cut int8) []Style {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(array); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		array[n-1], array[randIndex] = array[randIndex], array[n-1]
	}

	if cut != 0 {
		array = array[0:cut]
	}

	return array
}


func ShuffleAndCutStrings(array []string, cut int) []string {
	if cut > len(array) {
		cut = len(array)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(array); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		array[n-1], array[randIndex] = array[randIndex], array[n-1]
	}

	if cut != 0 {
		array = array[0:cut]
	}

	return array
}

func GetStyle(c *gin.Context) {
	db = include.GetDB()
	id := c.Params.ByName("id")

	var style Style
	var category Category

	if err := db.Select("style_id, name, impression").Where("style_id = ? ", id).First(&style).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Model(&style).Related(&category)

		style.Category = category

		c.JSON(200, style)
	}
}

func GetStyleRange(c *gin.Context){
	db = include.GetDB()
	var style Style

	if err := db.Where("beer_exam = true").Order("random()").First(&style).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Model(&style)

		c.JSON(200, style)
	}
}


func GetStyleQuestion(c *gin.Context){
	db = include.GetDB()

  variants := []string{"aroma", "appearance", "flavor", "mouthfeel"}
  n := rand.Int() % len(variants)
  variant := variants[n]

  var style Style
  var styles []Style
  var opt string
  var options []string
	var q Question
  var qs string

	rq := randBool() // reverse question


	if err := db.Preload("SimilarStyles").Where("beer_exam = true").Order("random()").First(&style).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Model(&style)

		styles = ShuffleAndCut(style.SimilarStyles, 3)
		style.SimilarStyles = nil

		a := getRand(4) // answer
		styles = Shuffle(styles) // shuffle current array
		o := insert(styles, a, style) // shuffled and organized options


		if rq {

			var rst string

			switch v := variant; v {
			case "aroma":
				rst = style.Aroma
			case "appearance":
				rst = style.Appearance
			case "flavor":
				rst = style.Flavor
			case "mouthfeel":
				rst = style.Mouthfeel
			default:
				rst = style.Impression
			}

			// trim period from the end.
			rst = TrimCharacter(rst, '.')

			// style characteristics array
			sca := strings.Split(rst, ". ")
			sca = ShuffleAndCutStrings(sca, 3)

			// style characteristics
			stcr := strings.Join(sca[:], ". ")

			qs = "Which beer style is better associated with the following " + variant + " characteristics: \"" + stcr + "\"?"
		} else {
			qs = "Which " + variant + " characteristics are better associated with " + style.Name + " (" + style.StyleId + ")?"
		}

		for i, l := 0, len(o); i < l; i++ {
			if rq {
				opt = o[i].Name + " (" + o[i].StyleId + ")"
			} else {
				switch v := variant; v {
				case "aroma":
					opt = o[i].Aroma
				case "appearance":
					opt = o[i].Appearance
				case "flavor":
					opt = o[i].Flavor
				case "mouthfeel":
					opt = o[i].Mouthfeel
				default:
					opt = o[i].Impression
				}
			}

			// trim period from the end
			opt = TrimCharacter(opt, '.')

			s := strings.Split(opt, ". ")
			s = ShuffleAndCutStrings(s, 3)

			strOpt := strings.Join(s[:], ". ") + ". "
			options = append(options, strOpt)
		}

		q.Options = options
		q.Question = qs
		q.Answer = a
		q.Exam = "beer"
		q.Topic = "styles"



		c.JSON(200, q)
	}
}
