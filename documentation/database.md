# Database

<!-- vscode-markdown-toc -->
* [Rules](#Rules)
* [DataSources](#DataSources)
* [Locales](#Locales)
* [Tags](#Tags)
* [Articles](#Articles)

<!-- vscode-markdown-toc-config
	numbering=false
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

## <a name='Rules'></a>Rules

| Field          | Type        |                          |
|----------------|-------------|--------------------------|
| id             | Serial      | Auto increment           |
| data_source_id | Int         | Reference to DataSources |
| url_pattern    | Text        |                          |
| name           | Varchar(64) |                          |
| created_at     | Datetime    |                          |
| updated_at     | Datetime    |                          |

## <a name='DataSources'></a>DataSources


| Field      | Type     |                    |
|------------|----------|--------------------|
| id         | Serial   | Auto increment     |
| name       | Text     |                    |
| url        | Text     | Restful compliance |
| created_at | Datetime |                    |
| updated_at | Datetime |                    |


## <a name='Locales'></a>Locales

| Field        | Type        |                |
|--------------|-------------|----------------|
| id           | Serial      | Auto increment |
| lang_code    | Varchar(32) |                |
| country_code | Varchar(32) |                |
| created_at   | Datetime    |                |
| updated_at   | Datetime    |                |

## <a name='Tags'></a>Tags

| Field      | Type     |                       |
|------------|----------|-----------------------|
| id         | Serial   | Auto increment        |
| rule_id    | Int      | Reference to Rules    |
| locale_id  | Int      | Reference to Language |
| type       | Text     |                       |
| attributes | json     |                       |
| value      | Text     |                       |
| created_at | Datetime |                       |
| updated_at | Datetime |                       |

## <a name='Articles'></a>Articles

| Field      | Type     |                       |
|------------|----------|-----------------------|
| id         | Serial   | Auto increment        |
| rule_id    | Int      | Reference to Rules    |
| locale_id  | Int      | Reference to Language |
| title      | Text     |                       |
| content    | Text     | Rich text format      |
| created_at | Datetime |                       |
| updated_at | Datetime |                       |
