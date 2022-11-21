package A

import "github.com/vbnetvbnet/B"

const version = "1.1"


func Version() string {
    return "A " + version + "\n" + B.Version()
}

