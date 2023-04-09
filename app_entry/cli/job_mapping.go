package cli

func getJobMapping() (jobMapping map[string]JobRunCmd) {
	jobCollection := initJobsCollection()

	jobMapping = map[string]JobRunCmd{
		"print_upper": jobCollection.jobsExample.PrintUpper,
		"check_panic": jobCollection.jobsExample.CheckPanicRecovery,
	}

	return
}
