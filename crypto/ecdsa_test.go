package crypto

import (
	"testing"
)

var pubkey1 = "0x0478897d20450edc12c8f9c5a5321b4dff563b441abe26bdeda03df243aaba127ce3bc843b54b39c2d48d1432f70d1edba3062e2513e2dd427f2c116a8f1c773bc"

func Test_HexToPublicKey(t *testing.T) {
	t.Logf("pubKey:%+v", HexToPublicKey(pubkey1))
}
