package watcher

import (
	"path/filepath"
	"time"

	"github.com/devlopersabbir/terrorserver/internal/logger"
	"github.com/fsnotify/fsnotify"
)

// Watch monitors a file path for changes and calls onChange whenever it is
// written or renamed. It blocks until the done channel is closed.
func Watch(path string, done <-chan struct{}, onChange func()) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	name := filepath.Clean(path)
	if err := w.Add(dir); err != nil {
		w.Close()
		return err
	}

	go func() {
		defer w.Close()
		var lastReload time.Time
		for {
			select {
			case <-done:
				return

			case event, ok := <-w.Events:
				if !ok {
					return
				}
				if filepath.Clean(event.Name) != name {
					continue
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) || event.Has(fsnotify.Chmod) {
					if time.Since(lastReload) < 250*time.Millisecond {
						continue
					}
					lastReload = time.Now()
					logger.Info("config changed (%s), reloading...", event.Name)
					onChange()
				}

			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				logger.Warn("watcher error: %v", err)
			}
		}
	}()

	return nil
}
