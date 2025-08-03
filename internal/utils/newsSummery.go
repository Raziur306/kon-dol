package utils

import (
	"github.com/Raziur306/kon-dol/internal/model"
	"strings"
)

func ExtractSummaryContent(story model.Story) string {
	var sb strings.Builder

	// Add headline and subheadline
	sb.WriteString(story.Headline + "\n")
	if story.Subheadline != "" {
		sb.WriteString(story.Subheadline + "\n")
	}

	// Add excerpt if available
	if story.Metadata.Excerpt != "" {
		sb.WriteString(story.Metadata.Excerpt + "\n")
	}

	// Extract main text content
	for _, card := range story.Cards {
		for _, element := range card.StoryElements {
			if element.Type == "text" && element.Text != "" {
				sb.WriteString(element.Text + "\n")
			}
		}
	}

	return sb.String()
}
