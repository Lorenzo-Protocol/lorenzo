UPGRADE_NAME ?= v2.0

propose-upgrade:
	UPGRADE_NAME=$(UPGRADE_NAME) ./proposal.sh upgrade

propose-general:
	@if [ -z "$(PROPOSAL_FILE)" ]; then \
		echo "Error: 'PROPOSAL_FILE' parameter is required for propose-general"; \
		echo "Usage: make propose-general PROPOSAL_FILE=<proposal_file_path>"; \
		exit 1; \
	fi
	@if [ ! -f "$(PROPOSAL_FILE)" ]; then \
		echo "Error: File '$(PROPOSAL_FILE)' does not exist"; \
		exit 1; \
	fi
	./proposal.sh general $(PROPOSAL_FILE)

.PHONY: propose-upgrade propose-general