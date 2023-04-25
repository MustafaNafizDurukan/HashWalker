package csv

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/MustafaNafizDurukan/HashWalker/pkg/constants"
)

func Save(filename string, fileHashesAny any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	switch fileHashes := fileHashesAny.(type) {
	case constants.MapPathHash:
		err = writer.Write([]string{"Path", "MD5"})
		if err != nil {
			return err
		}

		for path, hash := range fileHashes {
			err = writer.Write([]string{path, hash})
			if err != nil {
				return err
			}
		}
	case constants.MapEntryPathHash:
		err = writer.Write([]string{"Type", "Path", "MD5"})
		if err != nil {
			return err
		}

		for entryType, pathHash := range fileHashes {
			for path, hash := range pathHash {
				err = writer.Write([]string{string(entryType), path, hash})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func Read(filename string) (constants.MapPathHash, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	fileHashes := make(constants.MapPathHash)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		fileHashes[record[0]] = record[1]
	}

	return fileHashes, nil
}

func Compare(baseline, fileHashes constants.MapPathHash) constants.MapEntryPathHash {
	modifiedEntries := make(constants.MapEntryPathHash)

	append := func(entryType constants.EntryType, source constants.MapPathHash) {
		for path, hash := range source {
			if _, ok := modifiedEntries[entryType]; !ok {
				modifiedEntries[entryType] = constants.MapPathHash{path: hash}
				continue
			}

			modifiedEntries[entryType][path] = hash
		}
	}

	for path, bHash := range baseline {
		hash, ok := fileHashes[path]
		if !ok {
			append(constants.Deleted, constants.MapPathHash{path: bHash})

			continue
		}

		if bHash != hash {
			append(constants.ModifiedNew, constants.MapPathHash{path: hash})

			append(constants.ModifiedOld, constants.MapPathHash{path: bHash})

			delete(fileHashes, path)
			continue
		}

		delete(fileHashes, path)
	}

	for path, hash := range fileHashes {
		append(constants.Added, constants.MapPathHash{path: hash})
	}

	return modifiedEntries
}
