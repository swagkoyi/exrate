package list

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Save(currencies []string) error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "exrate", "currencies_list.json")

	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	data, err := json.Marshal(currencies)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load() ([]string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dir, "exrate", "currencies_list.json")

	var listz []string
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		} else {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &listz)
	if err != nil {
		return nil, err
	}
	return listz, nil
}
