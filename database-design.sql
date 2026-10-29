CREATE TABLE dbo.people (
	id uniqueidentifier DEFAULT newid() NOT NULL,
	first_name varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	last_name varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	date_of_birth date NULL,
	address varchar(500) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	created_at datetime2 DEFAULT sysdatetime() NOT NULL,
	updated_at datetime2 NULL,
	CONSTRAINT PK_people PRIMARY KEY (id)
);

INSERT INTO master.dbo.people
(id, first_name, last_name, date_of_birth, address, created_at, updated_at)
VALUES(N'B0EC7B00-BF24-4D8A-B058-46622B0393D6', N'test1', N'test1', '1990-07-01', N'bankok', '2026-10-28 17:58:33.857', NULL);
INSERT INTO master.dbo.people
(id, first_name, last_name, date_of_birth, address, created_at, updated_at)
VALUES(N'1F335ABD-504A-48F6-95E9-764F32836DAD', N'test2', N'test2', '2024-08-13', N'kanchanaburi', '2026-10-29 03:36:01.878', NULL);
INSERT INTO master.dbo.people
(id, first_name, last_name, date_of_birth, address, created_at, updated_at)
VALUES(N'40FAC5DD-1508-499A-81BE-B0FDF7558B09', N'test', N'test123', '1990-07-22', N'bangkok', '2026-10-28 18:43:42.898', NULL);