package internal

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)

type FileChanges []*FileChange

const (
	// RenameThreshold specifies the percentage of removed lines that
	// still exist in destination to consider them linked.
	RenameThreshold = 40
)

type changedFile struct {
	filepath string
	content  string
}

func NewFileChanges(changes object.Changes) (FileChanges, error) {
	newFiles := make([]changedFile, 0)
	deletedFiles := make([]changedFile, 0)

	fcs := make(FileChanges, 0)

	// Extracts raw addition/deletion stats and looks for any new/deleted files
	for _, c := range changes {
		patch, err := c.Patch()
		if err != nil {
			return nil, err
		}

		for _, fp := range patch.FilePatches() {
			// ignore empty patches (binary files, submodule refs updates)
			if len(fp.Chunks()) == 0 {
				continue
			}

			from, to := fp.Files()
			name := ""
			if from == nil {
				name = to.Path()
				newFiles = append(newFiles, changedFile{
					filepath: to.Path(),
					content:  fp.Chunks()[0].Content(),
				})
			} else if to == nil {
				name = from.Path()
				deletedFiles = append(deletedFiles, changedFile{
					filepath: from.Path(),
					content:  fp.Chunks()[0].Content(),
				})
			} else {
				name = from.Path()
			}
			fcs = append(fcs, NewFileChange(name, fp))
		}
	}

	renames := findRenames(newFiles, deletedFiles)
	for _, rename := range renames {
		for _, file := range fcs {
			if file.Name == rename.RenameFrom {
				file.Additions = 0
				file.Deletions = 0
				file.RenameTo = rename.RenameTo
				file.Similarity = rename.Similarity
			} else if file.Name == rename.Name {
				file.Additions = rename.Additions
				file.Deletions = rename.Deletions
				file.RenameFrom = rename.RenameFrom
				file.Similarity = rename.Similarity
			}
		}
	}

	return fcs, nil
}

func findRenames(newFiles, deletedFiles []changedFile) FileChanges {
	renames := make(FileChanges, 0)

OUTER:
	for _, deletedFile := range deletedFiles {
		// First try to Find identical matches. This efficiently limits the
		// number of comparisons later.
		for i, newFile := range newFiles {
			if deletedFile.content == newFile.content {

				// Delete file from array so we dont' process it twice
				length := len(newFiles)
				newFiles[length-1], newFiles[i] = newFiles[i], newFiles[length-1]
				newFiles = newFiles[:length-1]

				renames = append(renames, &FileChange{
					Name:       newFile.filepath,
					IsBinary:   false,
					Additions:  0,
					Deletions:  0,
					RenameTo:   newFile.filepath,
					RenameFrom: deletedFile.filepath,
					Similarity: 100.0,
				})
				continue OUTER
			}
		}

		// Otherwise start comparing all lines
		deletedLines := splitLines(deletedFile.content)
		deletedLinesCount := len(deletedLines)
		var highestMatch *FileChange
		for _, newFile := range newFiles {
			similarLines := 0
			addedLines := splitLines(newFile.content)
			addedLinesCount := len(addedLines)
			for _, deletedLine := range deletedLines {
				for i, addedLine := range addedLines {
					if addedLine == deletedLine {
						addedLines = append(addedLines[:i], addedLines[i+1:]...)
						similarLines += 1
						break
					}
				}
			}

			similarity := float32(similarLines) / float32(maxInt(addedLinesCount, deletedLinesCount)) * 100.0

			if similarity > RenameThreshold && (highestMatch == nil || similarity > highestMatch.Similarity) {
				// TODO check if newFile is already being used in another rename.
				//  If so compare similarity %. When higher it's the new match and
				//  re-queue the deleted file. When lower skip.
				highestMatch = &FileChange{
					Name:       newFile.filepath,
					IsBinary:   false,
					Additions:  addedLinesCount - similarLines,
					Deletions:  deletedLinesCount - similarLines,
					RenameTo:   newFile.filepath,
					RenameFrom: deletedFile.filepath,
					Similarity: similarity,
				}
			}
		}

		if highestMatch != nil {
			renames = append(renames, highestMatch)
		}
	}
	return renames
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func splitLines(content string) []string {
	return strings.Split(strings.Replace(content, "\r\n", "\n", -1), "\n")
}
