package lock
import (
    "errors"
)
// The reason for locking the issue or pull request conversation. Lock will fail if you don't use one of these reasons:   * `off-topic`   * `too heated`   * `resolved`   * `spam`
type LockPutRequestBody_lock_reason int

const (
    OFFTOPIC_LOCKPUTREQUESTBODY_LOCK_REASON LockPutRequestBody_lock_reason = iota
    TOOHEATED_LOCKPUTREQUESTBODY_LOCK_REASON
    RESOLVED_LOCKPUTREQUESTBODY_LOCK_REASON
    SPAM_LOCKPUTREQUESTBODY_LOCK_REASON
)

func (i LockPutRequestBody_lock_reason) String() string {
    return []string{"offTopic", "tooHeated", "resolved", "spam"}[i]
}
func ParseLockPutRequestBody_lock_reason(v string) (any, error) {
    result := OFFTOPIC_LOCKPUTREQUESTBODY_LOCK_REASON
    switch v {
        case "offTopic":
            result = OFFTOPIC_LOCKPUTREQUESTBODY_LOCK_REASON
        case "tooHeated":
            result = TOOHEATED_LOCKPUTREQUESTBODY_LOCK_REASON
        case "resolved":
            result = RESOLVED_LOCKPUTREQUESTBODY_LOCK_REASON
        case "spam":
            result = SPAM_LOCKPUTREQUESTBODY_LOCK_REASON
        default:
            return 0, errors.New("Unknown LockPutRequestBody_lock_reason value: " + v)
    }
    return &result, nil
}
func SerializeLockPutRequestBody_lock_reason(values []LockPutRequestBody_lock_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
