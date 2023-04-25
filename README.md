# HashWalker

HashWalker is a simple command-line tool that recursively walks through a directory, calculates the MD5 hash of each file, and saves the results to a CSV file. It can also compare two CSV files generated by the tool to detect new, modified, or deleted files.

## **Features**

- Walk through directories and subdirectories recursively
- Calculate MD5 hashes for each file
- Save file hashes to a CSV file
- Compare two CSV files to detect file changes

## **Usage**

HashWalker accepts three command-line arguments:

- **`-dir`**: The directory you want to walk through and hash. When this flag is used, the tool will calculate file hashes and save them to a CSV file named "file_hashes.csv" in the current working directory.
- **`-baseline`**: The baseline CSV file for comparison. This file should contain file hashes generated by HashWalker.
- **`-csv`**: The second CSV file for comparison. This file should also contain file hashes generated by HashWalker.

To generate a CSV file with file hashes:

```
./hashwalker --dir=/path/to/directory
```

To compare two CSV files:

```
./hashwalker --baseline=baseline.csv --csv=comparison.csv
```

When comparing CSV files, HashWalker will print new, modified, or deleted files to the console.

## **Installation**

To install and build HashWalker, make sure you have **[Go](https://golang.org/doc/install)** installed, and then run:

```
git clone https://github.com/yourusername/hashwalker.git
cd hashwalker
go build
```

This will create a binary named **`hashwalker`** in the current directory. You can move this binary to a directory in your **`PATH`** for easier access.

## **License**

This project is released under the **[MIT License](https://chat.openai.com/c/LICENSE)**.