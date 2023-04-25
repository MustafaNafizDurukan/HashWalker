package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/MustafaNafizDurukan/HashWalker/pkg/constants"
)

func Walk(root string) (constants.MapPathHash, error) {
	fileHashes := make(constants.MapPathHash)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		hash, err := hashFile(path)
		if err != nil {
			return nil
		}

		fileHashes[path] = hash

		return nil
	})

	return fileHashes, err
}

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
