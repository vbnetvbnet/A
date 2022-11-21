package A

import "github.com/vbnetvbnet/B"

const version = "1.2"


func Version() string {
    return "A " + version + "\n" + B.Version()
}

