.SILENT:

COVER_OPT="-cover"
TEST_OPTS=$(EXTRA_TEST_OPTS) $(COVER_OPT)

test: 
	go test ./... $(TEST_OPTS)
