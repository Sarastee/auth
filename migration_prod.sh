#!/bin/bash
source config/prod.env


sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v