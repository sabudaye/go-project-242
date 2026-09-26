package code

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := size(path, recursive, all)
	if err != nil {
		return "", err
	}

	return format(size, human), nil
}

func size(path string, recursive bool, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64

	for _, entry := range entries {
		if !all && isHidden(entry) {
			continue
		}

		if entry.IsDir() {
			if !recursive {
				continue
			}

			entryPath := filepath.Join(path, entry.Name())
			size, err := size(entryPath, recursive, all)
			if err != nil {
				log.Printf("%v %v", path, err)
				continue
			}

			total += size
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			log.Printf("%v %v", entry.Name(), err)
			continue
		}

		total += entryInfo.Size()
	}

	return total, nil
}

func isHidden(entry os.DirEntry) bool {
	return strings.HasPrefix(entry.Name(), ".")
}

func format(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%d", size)
	}

	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
		PB = TB * 1024
		EB = PB * 1024
	)

	switch {
	case size >= EB:
		return fmt.Sprintf("%.1fE", float64(size)/EB)
	case size >= PB:
		return fmt.Sprintf("%.1fP", float64(size)/PB)
	case size >= TB:
		return fmt.Sprintf("%.1fT", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.1fG", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.1fM", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.1fK", float64(size)/KB)
	default:
		return fmt.Sprintf("%dB", size)
	}
}
