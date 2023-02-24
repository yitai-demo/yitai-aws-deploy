package model

type GetActiveSchedulerJobsRsp struct {
	ErrorCode     int             `json:"errcode"`
	ErrorMessage  string          `json:"errmsg"`
	SchedulerJobs []*SchedulerJob `json:"data"`
}

type SchedulerJob struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	SchedulerType       string `json:"schedulerType"`
	SchedulerDefinition string `json:"schedulerDefinition"`
	Action              string `json:"action"`
}
