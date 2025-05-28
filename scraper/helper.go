package scraper

import (
	"net/url"
	"path"
)

// extractSlug takes a raw URL string and returns the last path segment (slug).
// It assumes the URL is valid and a slug exists; errors from url.Parse are ignored for brevity.
func extractSlug(rawURL string) string {
	parsedURL, _ := url.Parse(rawURL) // Parse the raw URL string
	return path.Base(parsedURL.Path)  // Get the last element of the path
}
