package utils

import (
	"strings"

	"github.com/Raziur306/kon-dol/internal/model"
)

func IsPotentialPoliticalViolence(articleText string, tags []model.Tag) bool {
	politicalParties := []string{
		"বিএনপি", "আওয়ামী লীগ", "জাতীয় পার্টি", "জামায়াত", "জাতীয় সমাজতান্ত্রিক দল", "নগর উন্নয়ন পার্টি", "ইসলামী আন্দোলন", "বাসদ", "ন্যাপ", "কমিউনিস্ট পার্টি", "নতুন ধারা", "গণফোরাম", "নাগরিক ঐক্য", "বাম মোর্চা", "লেবার পার্টি", "আমরা জনগণ পার্টি", "নিপ", "জাতীয় পার্টি (জেপি)", "এনসিপির", "জাতীয় নাগরিক পার্টি",
	}

	violenceKeywords := []string{
		"চাঁদা", "হত্যা", "খুন", "ধর্ষণ", "বোমা", "গ্রেনেড", "সংঘর্ষ", "গ্রেপ্তার", "ধাওয়া", "আহত", "মারধর", "জখম", "নিহত", "অস্ত্র", "গুলিবর্ষণ", "পেটানো", "হুমকি",
	}

	hasParty := false
	hasViolence := false

	combinedText := articleText + " " + tagNamesToString(tags) // Merge both fields for easier checking

	for _, name := range politicalParties {
		if strings.Contains(combinedText, name) {
			hasParty = true
			break
		}
	}
	for _, word := range violenceKeywords {
		if strings.Contains(combinedText, word) {
			hasViolence = true
			break
		}
	}

	return hasParty && hasViolence
}

func tagNamesToString(tags []model.Tag) string {
	var names []string
	for _, tag := range tags {
		names = append(names, tag.Name)
	}
	return strings.Join(names, " ")
}
