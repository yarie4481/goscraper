package webscraper

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/chromedp/chromedp"
)

// Scraper holds configuration for the web scraper
type Scraper struct {
	Headless       bool
	TimeoutSeconds int
	LogOutput      io.Writer
	logger         *log.Logger
}

// New creates a new Scraper instance with default configuration
func New() *Scraper {
	return &Scraper{
		Headless:       true,
		TimeoutSeconds: 120,
		LogOutput:      os.Stdout,
	}
}

// ScrapeText scrapes visible text from a webpage
func (s *Scraper) ScrapeText(url string) (string, error) {
	// Initialize logger
	s.logger = log.New(s.LogOutput, "[webscraper] ", log.LstdFlags)
	s.logger.Println("🚀 Starting browser automation...")

	// 1. Find Chrome
	chromePath, err := findChromePath()
	if err != nil {
		return "", fmt.Errorf("chrome not found: %v", err)
	}
	s.logger.Printf("🔍 Using Chrome at: %s", chromePath)

	// 2. Configure browser options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(chromePath),
		chromedp.Flag("headless", s.Headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
	)

	// 3. Create browser context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(s.logger.Printf),
		chromedp.WithDebugf(s.logger.Printf),
	)
	defer cancel()

	// 4. Set timeout
	ctx, cancel = context.WithTimeout(ctx, time.Duration(s.TimeoutSeconds)*time.Second)
	defer cancel()

	// 5. Run scraping tasks
	s.logger.Println("🌐 Navigating to URL:", url)
	var text string
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			s.logger.Println("🕒 Waiting for page to load...")
			return nil
		}),
		chromedp.Sleep(5*time.Second),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			s.logger.Println("📄 Page loaded, extracting text...")
			return nil
		}),
		chromedp.Text("body", &text),
		chromedp.Sleep(2*time.Second),
	)

	if err != nil {
		return "", fmt.Errorf("browser operation failed: %v", err)
	}

	s.logger.Printf("📝 Extracted %d characters", len(text))
	return text, nil
}

// findChromePath searches for Chrome executable in common locations
func findChromePath() (string, error) {
	paths := []string{
		`C:\Program Files\Google\Chrome\Application\chrome.exe`,
		`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
		`C:\Users\*\AppData\Local\Google\Chrome\Application\chrome.exe`,
		"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		"/usr/bin/google-chrome",
		"/usr/bin/chromium",
		"/usr/bin/chromium-browser",
	}

	// Check common executable names
	for _, name := range []string{"chrome", "google-chrome", "chromium"} {
		if path, err := exec.LookPath(name); err == nil {
			return path, nil
		}
	}

	// Check specific paths
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("Chrome not found in standard locations")
}