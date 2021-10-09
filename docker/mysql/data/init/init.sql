CREATE DATABASE IF NOT EXISTS short_cut_master;
USE short_cut_master;

CREATE TABLE IF NOT EXISTS users (
  id             INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name           VARCHAR(256) NOT NULL,
  is_admin       BOOLEAN NOT NULL DEFAULT false,
  google_user_id VARCHAR(256),
  created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS quizzes (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name        VARCHAR(256) NOT NULL,
  type        ENUM("macOS", "windows"),
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS rankings (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  quiz_id     INT NOT NULL,
  ranking     INT NOT NULL DEFAULT 0,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (quiz_id)
    REFERENCES quizzes(id)
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS user_rankings (
  id             INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  user_id        INT NOT NULL,
  ranking_id     INT NOT NULL,
  created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (ranking_id)
    REFERENCES rankings(id)
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS questions (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  quiz_id     INT NOT NULL,
  contents    VARCHAR(256) NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (quiz_id)
    REFERENCES quizzes(id)
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS answers (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  question_id INT NOT NULL,
  contents    VARCHAR(256) NOT NULL,
  is_correct  BOOLEAN NOT NULL DEFAULT false,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (question_id)
    REFERENCES questions(id)
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8;

INSERT INTO users (name, is_admin) VALUES ("テストユーザー1", false),("テストユーザー2", false),("テストユーザー3", false), ("テストユーザー4", false), ("テスト管理者", true);

INSERT INTO quizzes (name, type) VALUES ("Slack", "macOS"), ("VSCode", "macOS"), ("Chrome", "macOS"), ("Github", "macOS");

INSERT INTO rankings (quiz_id, ranking) VALUES (1, 1), (2, 1), (3, 1), (4, 1);

INSERT INTO user_rankings (ranking_id, user_id) VALUES (1, 1), (2, 1), (3, 3), (4, 4);

INSERT INTO questions (quiz_id, contents) VALUES (1, "DMを閲覧するには？");

INSERT INTO answers (question_id, contents, is_correct) VALUES (1, "⌘+shit+↓", false), (1, "⌘+shit+K", true), (1, "⌘+K", false), (1, "⌘+T", false);
