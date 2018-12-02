-- 事前作業：ユーザ／データベースの作成、データベース切替
-- スーパユーザでログイン (> psql -U postgres)
-- # CREATE ROLE denki_user WITH LOGIN PASSWORD 'denki_user';
-- # \q
-- > psql -U denki_user -d postgres
-- $ \q
-- > psql -U postgres
-- # ALTER ROLE denki_user SUPERUSER; (後で権限変更予定)
-- #  # ALTER ROLE denki_user LOGIN;
-- # \q
-- > psql -U denki_user -d postgres
-- # CREATE DATABASE denki;
-- # \c denki

-- テーブルの作成
-- \i [[filepath]]
CREATE TABLE SUPER_VISOR (
	super_visor_id VARCHAR(30) PRIMARY KEY,
	super_visor_password VARCHAR(20) NOT NULL
);

CREATE TABLE CLASS (
	class_id BIGINT PRIMARY KEY,
	class_name VARCHAR(100) NOT NULL
);

CREATE TABLE MAKER (
	maker_id BIGINT PRIMARY KEY,
	maker_name VARCHAR(100) NOT NULL
);

CREATE TABLE GOODS (
	goods_id BIGINT PRIMARY KEY,
	goods_name VARCHAR(100) NOT NULL,
	class_id INTEGER NOT NULL REFERENCES CLASS(CLASS_ID),
	maker_id INTEGER NOT NULL REFERENCES MAKER(MAKER_ID),
	model_number VARCHAR(30) NOT NULL,
	specs TEXT,
	indicated_price INTEGER,
	purchase_price NUMERIC,
	stock INTEGER,
	is_deleted BOOLEAN NOT NULL,
	update_super_visor_id VARCHAR(30) NOT NULL,
	update_date TIMESTAMP NOT NULL,
	update_version_id BIGINT NOT NULL
);
