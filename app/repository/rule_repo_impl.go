package repository

import "database/sql"

func scanRule(rows *sql.Rows) (*Rule, error) {
	var rule Rule
	var err error
	if err = rows.Scan(&rule.ID, &rule.Name, &rule.UrlPattern, &rule.UpdatedAt, &rule.CreatedAt); err != nil {
		return nil, err
	}
	return &rule, nil
}
