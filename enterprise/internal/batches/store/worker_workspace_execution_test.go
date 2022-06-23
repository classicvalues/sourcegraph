	"github.com/cockroachdb/errors"
	db := dbtest.NewDB(t)

	workspace := &btypes.BatchSpecWorkspace{BatchSpecID: batchSpec.ID, RepoID: repo.ID, Steps: []batcheslib.Step{}}
	if err := s.CreateBatchSpecWorkspace(ctx, workspace); err != nil {
		t.Fatal(err)
	}

	job := &btypes.BatchSpecWorkspaceExecutionJob{BatchSpecWorkspaceID: workspace.ID}
	if err := ct.CreateBatchSpecWorkspaceExecutionJob(ctx, s, ScanBatchSpecWorkspaceExecutionJob, job); err != nil {
		t.Fatal(err)
	}

		"Nsw12JxoLSHN4ta6D3G7FQ",
stdout: {"operation":"CACHE_RESULT","timestamp":"2021-11-04T12:43:19.551Z","status":"SUCCESS","metadata":{"key":"Nsw12JxoLSHN4ta6D3G7FQ","value":{"diff":"diff --git README.md README.md\nindex 1914491..d6782d3 100644\n--- README.md\n+++ README.md\n@@ -3,4 +3,7 @@ This repository is used to test opening and closing pull request with Automation\n \n (c) Copyright Sourcegraph 2013-2020.\n (c) Copyright Sourcegraph 2013-2020.\n-(c) Copyright Sourcegraph 2013-2020.\n\\ No newline at end of file\n+(c) Copyright Sourcegraph 2013-2020.this is step 2\n+this is step 3\n+this is step 4\n+previous_step.modified_files=[README.md]\ndiff --git README.txt README.txt\nnew file mode 100644\nindex 0000000..888e1ec\n--- /dev/null\n+++ README.txt\n@@ -0,0 +1 @@\n+this is step 1\ndiff --git my-output.txt my-output.txt\nnew file mode 100644\nindex 0000000..257ae8e\n--- /dev/null\n+++ my-output.txt\n@@ -0,0 +1 @@\n+this is step 5\n","changedFiles":{"modified":["README.md"],"added":["README.txt","my-output.txt"],"deleted":null,"renamed":null},"outputs":{"myOutput":"my-output.txt"},"Path":""}}}
	_, err := workStore.AddExecutionLogEntry(ctx, int(job.ID), entry, dbworkerstore.ExecutionLogEntryOptions{})
	if err != nil {
		t.Fatal(err)

	executionStore := &batchSpecWorkspaceExecutionWorkerStore{Store: workStore, observationContext: &observation.TestContext}
	setProcessing := func(t *testing.T) {
	attachAccessToken := func(t *testing.T) int64 {
		tokenID, _, err := database.AccessTokens(db).CreateInternal(ctx, user.ID, []string{"user:all"}, "testing", user.ID)
			t.Fatal(err)
		if err := s.SetBatchSpecWorkspaceExecutionJobAccessToken(ctx, job.ID, tokenID); err != nil {
			t.Fatal(err)
		return tokenID
	assertJobState := func(t *testing.T, want btypes.BatchSpecWorkspaceExecutionJobState) {
		reloadedJob, err := s.GetBatchSpecWorkspaceExecutionJob(ctx, GetBatchSpecWorkspaceExecutionJobOpts{ID: job.ID})
			t.Fatalf("failed to reload job: %s", err)
		if have := reloadedJob.State; have != want {
			t.Fatalf("wrong job state: want=%s, have=%s", want, have)
		setProcessing(t)
		tokenID := attachAccessToken(t)
		assertJobState(t, btypes.BatchSpecWorkspaceExecutionJobStateCompleted)
			entries, err := s.ListBatchSpecExecutionCacheEntries(ctx, ListBatchSpecExecutionCacheEntriesOpts{Keys: []string{wantKey}})
			var cachedExecutionResult *execution.Result
		_, err = database.AccessTokens(db).GetByID(ctx, tokenID)
		if err != database.ErrAccessTokenNotFound {
			t.Fatalf("access token was not deleted")
		ok, err := executionStore.MarkComplete(context.Background(), int(job.ID), opts)
			t.Fatalf("MarkComplete failed. ok=%t, err=%s", ok, err)
		assertJobState(t, btypes.BatchSpecWorkspaceExecutionJobStateCompleted)
	t.Run("token set but deletion fails", func(t *testing.T) {
		tokenID := attachAccessToken(t)
		database.Mocks.AccessTokens.HardDeleteByID = func(id int64) error {
			if id != tokenID {
				t.Fatalf("wrong token deleted")
			return errors.New("internal database error")
		defer func() { database.Mocks.AccessTokens.HardDeleteByID = nil }()
		ok, err := executionStore.MarkComplete(context.Background(), int(job.ID), opts)
		if !ok || err != nil {
			t.Fatalf("MarkComplete failed. ok=%t, err=%s", ok, err)
		assertJobState(t, btypes.BatchSpecWorkspaceExecutionJobStateFailed)
	})