package watcher

import (
	"github.com/fsnotify/fsnotify"
	"github.com/devlopersabbir/terrorserver/internal/logger"
)

// Watch monitors a file path for changes and calls onChange whenever it is
// written or renamed. It blocks until the done channel is closed.
func Watch(path string, done <-chan struct{}, onChange func()) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	if err := w.Add(path); err != nil {
		w.Close()
		return err
	}

	go func() {
		defer w.Close()
		for {
			select {
			case <-done:
				return

			case event, ok := <-w.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
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
