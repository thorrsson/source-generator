package models
import (
    "errors"
)
// The phase of the lifecycle that the check is currently in.
type CheckRun_status int

const (
    QUEUED_CHECKRUN_STATUS CheckRun_status = iota
    IN_PROGRESS_CHECKRUN_STATUS
    COMPLETED_CHECKRUN_STATUS
)

func (i CheckRun_status) String() string {
    return []string{"queued", "in_progress", "completed"}[i]
}
func ParseCheckRun_status(v string) (any, error) {
    result := QUEUED_CHECKRUN_STATUS
    switch v {
        case "queued":
            result = QUEUED_CHECKRUN_STATUS
        case "in_progress":
            result = IN_PROGRESS_CHECKRUN_STATUS
        case "completed":
            result = COMPLETED_CHECKRUN_STATUS
        default:
            return 0, errors.New("Unknown CheckRun_status value: " + v)
    }
    return &result, nil
}
func SerializeCheckRun_status(values []CheckRun_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
