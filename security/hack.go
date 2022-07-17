package security

import (
	"fmt"
	"log"
)

func AlertAdulteration(cause string) {
	log.Panic(fmt.Errorf("DETECTED HACK: %v. EXITTING", cause))
}