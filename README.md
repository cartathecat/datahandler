## Data Handler

Loads port data from a static JSON file (ports.json)

The Port struct models the list of ports:

```
type Port struct {
	Name        string    `json:"name"`        // Name of port
	City        string    `json:"city"`        // City of port
	Country     string    `json:"country"`     // Country of port
	Alias       []string  `json:"alias"`       // Alias'
	Regions     []string  `json:"regions"`     // Regions
	Coordinates []float64 `json:"coordinates"` // Long/Lat coordinates
	Province    string    `json:"province"`    // Province
	TimeZone    string    `json:"timezone"`    // Time zone
	Unlocks     []string  `json:"unlocks"`     // Unlocks
	Code        string    `json:"code"`        // Code
}

```

The following functions are exposed:

Retrieve an individual port
```
func (key *PortID) GetPort() (Port, error)
```

Retrieve a list of port identifiers
```
func ListOfPorts() []PortID
```

Retrieve ALL port details
```
func ListAllPorts() map[string]interface{}
```

