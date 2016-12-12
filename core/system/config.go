package system

// /*
// LoadConfig returns a Config struct loaded from a specified
// path.
// */
// func LoadConfig(path string) (config.Config, error) {
// 	_, err := os.Stat(path)
// 	if os.IsNotExist(err) {
// 		return Config{}, ErrConfigDoesntExist
// 	}
// 	if err == nil {
// 		configFile, openErr := os.Open(path)
// 		if openErr != nil {
// 			return Config{}, err
// 		}
// 		decoder := json.NewDecoder(configFile)
// 		config := Config{}
// 		decoder.Decode(&config)
// 		return config, nil
// 	}
// 	return Config{}, err
// }
