package database

/*
Naming convention
1. small letters everywhere
2. "_" - separator for hierarchy levels or meaning
3. all starts small letters
4. as short as possible
5. some standard values
  - "t" - table
  - "v" - view
  - "m" - map table
6. use singular in name
EXAMPLE:
t_task - table task
v_trigger - view task
m_task_trigger - view that maps v_task with v_trigger - it's to get some results after

Some details
1. table used as storage but, you can not add/change/delete anything there directly
2. to do add/change/delete used view
3. mapping tables used to describe relationship (think about it like a graph where mapping are graph lines, and tables/view - vertices)
4. stored procedures help to do some cool work

*/

var initDB = `
-- disable foreign key constraint check
PRAGMA foreign_keys=off;
-- start a transaction
BEGIN TRANSACTION;

-- actions with t_task
-- create table for first time
CREATE TABLE IF NOT EXISTS t_task (
	id INTEGER PRIMARY KEY,
	command TEXT NOT NULL,
	datetime INTEGER NOT NULL
);
-- new structure
CREATE TABLE IF NOT EXISTS t_task_new (
	id INTEGER PRIMARY KEY,
	command TEXT NOT NULL,
	datetime INTEGER NOT NULL
);
-- copy data
INSERT INTO t_task_new (id, command, datetime)
SELECT id, command, datetime
FROM t_task;
-- drop the table
DROP TABLE t_task;
-- rename the new_table to the table
ALTER TABLE t_task_new RENAME TO t_task;

-- actions with t_trigger
-- create table for first time
CREATE TABLE IF NOT EXISTS t_trigger (
	id INTEGER PRIMARY KEY,
	condition TEXT NOT NULL,
	datetime INTEGER NOT NULL
);
-- new structure
CREATE TABLE IF NOT EXISTS t_trigger_new (
	id INTEGER PRIMARY KEY,
	condition TEXT NOT NULL,
	datetime INTEGER NOT NULL
);
-- copy data
INSERT INTO t_trigger_new (id, condition, datetime)
SELECT id, condition, datetime
FROM t_trigger;
-- drop the table
DROP TABLE t_trigger;
-- rename the new_table to the table
ALTER TABLE t_trigger_new RENAME TO t_trigger;

-- actions with m_task_trigger
-- create table for first time
CREATE TABLE IF NOT EXISTS m_task_trigger (
	task_id INTEGER,
	trigger_id INTEGER,
	PRIMARY KEY (task_id, trigger_id),
	FOREIGN KEY (task_id)
		REFERENCES t_task (id)
			ON DELETE NO ACTION
			ON UPDATE NO ACTION,
	FOREIGN KEY (trigger_id)
		REFERENCES t_trigger (id)
			ON DELETE NO ACTION
			ON UPDATE NO ACTION
);

-- commit the transaction
COMMIT;
-- enable foreign key constraint check
PRAGMA foreign_keys=on;
`

var addExampleData = `
INSERT INTO t_task (id, command, datetime)
	VALUES(0,'echo "I am the best!',strftime('%s','now'));
INSERT INTO t_task (id, command, datetime)
	VALUES(1,'echo "No, I am the best!"',strftime('%s','now'));
INSERT INTO t_task (id, command, datetime)
	VALUES(2,'echo "Way to the king! both of you!"',strftime('%s','now'));

INSERT INTO t_trigger (id, condition, datetime)
	VALUES(0,'echo "first"',strftime('%s','now'));
INSERT INTO t_trigger (id, condition, datetime)
	VALUES(1,'echo "second"',strftime('%s','now'));
INSERT INTO t_trigger (id, condition, datetime)
	VALUES(2,'echo "third"',strftime('%s','now'));

INSERT INTO m_task_trigger (task_id, trigger_id)
	VALUES(0,1);
INSERT INTO m_task_trigger (task_id, trigger_id)
	VALUES(1,2);
INSERT INTO m_task_trigger (task_id, trigger_id)
	VALUES(2,0);
`

var delExampleData = `
DELETE FROM t_task
	WHERE id = 0;
DELETE FROM t_task
	WHERE id = 1;
DELETE FROM t_task
	WHERE id = 2;

DELETE FROM t_trigger
	WHERE id = 0;
DELETE FROM t_trigger
	WHERE id = 1;
DELETE FROM t_trigger
	WHERE id = 2;

DELETE FROM m_task_trigger
	WHERE task_id = 0 and trigger_id = 1;
DELETE FROM m_task_trigger
	WHERE task_id = 1 and trigger_id = 2;
DELETE FROM m_task_trigger
	WHERE task_id = 2 and trigger_id = 0;
`
