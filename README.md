## This is a simple package that helps load envronmental variables from a yaml file using viper


## usage
```go
import (
	"fmt"
	"log"

	loadcfg "github.com/xtasysensei/go-yaml-config/loadCfg"
)

func main() {
	config, err := loadcfg.LoadYamlCfg()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Postgres Host: %s\n", config.Db.Psql.Host)
	fmt.Printf("Postgres Port: %s\n", config.Db.Psql.Port)
	fmt.Printf("Server Port: %s\n", config.SRV.Port)
}

```

And a toml file should be present in the root directory or you can modify the `loadcfg.go` to specify another directory

```yaml
database:
  postgresql:
    hostname: localhost
    port: 5432
    username: sensei
    password: 12345

  redis:
    hostname: shem

server:
  port: 8000
```

## License
Use as you please
