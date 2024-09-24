package json

import (
	"encoding/json"
	"fmt"
	"log"
	"opg-analysis/internal/trade"
	"os"
)

// deliverer is a struct that represents a JSON deliverer.
// It is responsible for writing trade selections to a JSON file.
type deliverer struct {
	filepath string
}

// Deliver writes the given trade selections to a JSON file.
// It returns an error if there was a problem creating the file or encoding the selections.
func (d *deliverer) Deliver(selections []trade.Selection) error {
	file, err := os.Create(d.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(selections)

	if err != nil {
		return fmt.Errorf("error encoding selections: %w", err)
	}

	log.Printf("Finished writing output to %s\n", d.filepath)
	return nil
}

// NewDeliverer creates a new JSON deliverer with the given file path.
func NewDeliverer(filePath string) trade.Deliverer {
	return &deliverer{
		filepath: filePath,
	}
}
