# traq-cname2id

traQのチャンネル名をチャンネルIDに変換するライブラリ

[go-traq](https://github.com/traPtitech/go-traq)と併用する前提です

```
go get github.com/traPtitech/go-traq
go get github.com/oribe1115/traq-cname2id
```

```go
package main

import (
	"context"
	"fmt"

    cname2id "github.com/oribe1115/traq-cname2id"
	traq "github.com/traPtitech/go-traq"
)

const TOKEN = "/* your token */"

func main() {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

    c := cname2id.NewConverter(client, auth)
    id, _ := c.GetChannelID("#a/bb/ccc")

    fmt.Println(id)
}
```
