migration:
	@last=$$(ls internal/migrations | grep -E '^[0-9]{6}_.+\.go$$' | sort | tail -n 1 | cut -d_ -f1); \
	if [ -z "$$last" ]; then \
		next=1; \
	else \
		next=$$((10#$$last + 1)); \
	fi; \
	num=$$(printf "%06d" $$next); \
	touch internal/migrations/$${num}_$(name).go; \
	echo "Created migrations/$${num}_$(name).go"
