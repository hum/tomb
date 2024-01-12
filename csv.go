package main

import (
	"encoding/csv"
	"os"
	"strings"
)

// CSVHandler handles everything regarding filesytem I/O.
type CSVHandler struct {
	f    *os.File
	cswW *csv.Writer
}

// NewCSVHandler allows the creation, and initialisation, of a new CSV handler for a given file.
// `headers` allows specifying the first row of the CSV file to be populated with labels.
//
// @TODO: refactor headers out of the init.
func NewCSVHandler(fp string, headers []string) (*CSVHandler, error) {
	file, err := os.Create(fp)
	if err != nil {
		return nil, err
	}

	w := csv.NewWriter(file)

	// Immediately write headers to the open csv file
	w.Write(headers)

	return &CSVHandler{
		f:    file,
		cswW: w,
	}, nil
}

// SaveMessagesToFile manages appending the normalised messages into the open CSV file.
// Only use when the `CSVHandler` was initialised with `NewCSVHandler` function.
func (h *CSVHandler) SaveMessagesToFile(m []*NormalisedDiscordMessage) {
	for _, msg := range m {
		// Replace new-line endings to not corrupt the CSV file
		content := strings.Replace(msg.Content, "\n", "", -1)

		h.cswW.Write([]string{msg.Timestamp.String(), msg.GuildId, msg.ChannelId, msg.Username, content})
	}
}

// Close asserts that the underlying file descriptor and csv reader are properly closed.
// Always call this function before exiting the program.
func (h *CSVHandler) Close() {
	h.f.Close()
	h.cswW.Flush()
}
