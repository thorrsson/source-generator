package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowRunUsage_billable_MACOS 
type WorkflowRunUsage_billable_MACOS struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The job_runs property
    job_runs []WorkflowRunUsage_billable_MACOS_job_runsable
    // The jobs property
    jobs *int32
    // The total_ms property
    total_ms *int32
}
// NewWorkflowRunUsage_billable_MACOS instantiates a new WorkflowRunUsage_billable_MACOS and sets the default values.
func NewWorkflowRunUsage_billable_MACOS()(*WorkflowRunUsage_billable_MACOS) {
    m := &WorkflowRunUsage_billable_MACOS{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkflowRunUsage_billable_MACOSFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowRunUsage_billable_MACOSFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowRunUsage_billable_MACOS(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkflowRunUsage_billable_MACOS) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkflowRunUsage_billable_MACOS) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["job_runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowRunUsage_billable_MACOS_job_runsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkflowRunUsage_billable_MACOS_job_runsable, len(val))
            for i, v := range val {
                res[i] = v.(WorkflowRunUsage_billable_MACOS_job_runsable)
            }
            m.SetJobRuns(res)
        }
        return nil
    }
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobs(val)
        }
        return nil
    }
    res["total_ms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMs(val)
        }
        return nil
    }
    return res
}
// GetJobRuns gets the job_runs property value. The job_runs property
func (m *WorkflowRunUsage_billable_MACOS) GetJobRuns()([]WorkflowRunUsage_billable_MACOS_job_runsable) {
    return m.job_runs
}
// GetJobs gets the jobs property value. The jobs property
func (m *WorkflowRunUsage_billable_MACOS) GetJobs()(*int32) {
    return m.jobs
}
// GetTotalMs gets the total_ms property value. The total_ms property
func (m *WorkflowRunUsage_billable_MACOS) GetTotalMs()(*int32) {
    return m.total_ms
}
// Serialize serializes information the current object
func (m *WorkflowRunUsage_billable_MACOS) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("jobs", m.GetJobs())
        if err != nil {
            return err
        }
    }
    if m.GetJobRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobRuns()))
        for i, v := range m.GetJobRuns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("job_runs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_ms", m.GetTotalMs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkflowRunUsage_billable_MACOS) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJobRuns sets the job_runs property value. The job_runs property
func (m *WorkflowRunUsage_billable_MACOS) SetJobRuns(value []WorkflowRunUsage_billable_MACOS_job_runsable)() {
    m.job_runs = value
}
// SetJobs sets the jobs property value. The jobs property
func (m *WorkflowRunUsage_billable_MACOS) SetJobs(value *int32)() {
    m.jobs = value
}
// SetTotalMs sets the total_ms property value. The total_ms property
func (m *WorkflowRunUsage_billable_MACOS) SetTotalMs(value *int32)() {
    m.total_ms = value
}
