package tomladapter

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/pelletier/go-toml"
)

func init() {
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

// 어댑터는 TOML 형식의 Caddy 구성을 Json으로 변환
type Adapter struct{}

// TOML형식의 바디를 Json으로 변환
func (a Adapter) Adapt(body []byte, _ map[string]interface{}) (
	[]byte, []caddyconfig.Warning, error) {

	tree, err := toml.LoadBytes(body)
	if err != nil {
		return nil, nil, err
	}

	b, err := json.Marshal(tree.ToMap())

	return b, nil, err
}
