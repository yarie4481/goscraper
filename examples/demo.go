package main

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

func main() {
	log.Println("🚀 Starting browser automation...")

	// Configure logging to file
	logFile, err := os.OpenFile("scraper.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("❌ Failed to create log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	text, err := scrapeTextFromBrowser("https://www.ashewatechnology.com/")
	if err != nil {
		log.Fatalf("❌ Failed: %v", err)
	}

	saveToFile("output.txt", text)
	log.Println("✅ Success! Extracted visible text.")
}

func scrapeTextFromBrowser(url string) (string, error) {
	// 1. Find Chrome
	chromePath, err := findChromePath()
	if err != nil {
		return "", fmt.Errorf("chrome not found: %v", err)
	}
	log.Printf("🔍 Using Chrome at: %s", chromePath)

	// 2. Configure browser options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(chromePath),
		chromedp.Flag("headless", false),                  // Visible browser
		chromedp.Flag("disable-gpu", true),                // Needed for Windows
		chromedp.Flag("no-sandbox", true),                 // Bypass OS security
		chromedp.Flag("disable-extensions", true),         // Disable extensions
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),                   // Set window size
	)

	// 3. Create browser context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(log.Printf),  // Enable browser logging
		chromedp.WithDebugf(log.Printf), // Enable debug logging
	)
	defer cancel()

	// 4. Set generous timeout (2 minutes)
	ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	// 5. Run scraping tasks
	log.Println("🌐 Navigating to URL:", url)
	var text string
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("🕒 Waiting for page to load...")
			return nil
		}),
		chromedp.Sleep(5*time.Second),                     // Initial wait
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("📄 Page loaded, extracting text...")
			return nil
		}),
		chromedp.Text("body", &text),
		chromedp.Sleep(2*time.Second),                     // Final wait
	)

	if err != nil {
		return "", fmt.Errorf("browser operation failed: %v", err)
	}

	log.Printf("📝 Extracted %d characters", len(text))
	return text, nil
}

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

func saveToFile(filename, content string) {
	if len(content) == 0 {
		log.Printf("⚠️ No content to save to %s", filename)
		return
	}

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		log.Printf("❌ Failed to save %s: %v", filename, err)
	} else {
		log.Printf("💾 Saved %d characters to %s", len(content), filename)
	}
}