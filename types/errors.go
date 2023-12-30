package types

import (
	"fmt"
	"strings"
)

type Error string

// This error is returned when the replica set is not initialized
// this error used in utils/is_init_replica_set.go to check if the replica set is initialized
const ErrorNoReplicaSetConfig Error = "no replset config has been received"

// Match checks if the error is equal to the error or contains the error
func (e Error) Match(err error) bool {
	return string(e) == err.Error() || strings.Contains(err.Error(), string(e))
}

var ErrorNotAddMember = fmt.Errorf("db not ok")
var ErrorNotRemoveMember = fmt.Errorf("db not ok")
