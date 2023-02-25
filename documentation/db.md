#### Events Table
| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | INT | PRIMARY KEY, AUTO_INCREMENT |
| name | VARCHAR(255) | NOT NULL |
| date | DATETIME | NOT NULL |
| location | VARCHAR(255) | NOT NULL |
| description | TEXT | NOT NULL |
| created_by | INT | NOT NULL |
| total_participant | INT | NOT NULL |
| created_at | DATETIME | NOT NULL |

#### Participants Table
| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | INT | PRIMARY KEY, AUTO_INCREMENT |
| name | VARCHAR(255) | NOT NULL |
| email | VARCHAR(255) | NOT NULL |
| event_id | INT | NOT NULL |
| created_at | DATETIME | NOT NULL |

#### Users Table
| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | INT | PRIMARY KEY, AUTO_INCREMENT |
| name | VARCHAR(255) | NOT NULL |
| email | VARCHAR(255) | NOT NULL |
| password | VARCHAR(255) | NOT NULL |
| created_at | DATETIME | NOT NULL |