CREATE DATABASE IF NOT EXISTS go_dev_api;
USE go_dev_api;

CREATE TABLE IF NOT EXISTS users (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name        VARCHAR(256) NOT NULL,
  is_admin    BOOLEAN NOT NULL DEFAULT false,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
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
  user_id     INT NOT NULL,
  ranking     INT NOT NULL DEFAULT 0,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (quiz_id)
    REFERENCES quizzes(id)
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

INSERT INTO users (id, name, is_admin) VALUES (1, "テストユーザー1", false),(2, "テストユーザー2", false),(3, "テストユーザー3", false), (4, "テストユーザー4", false), (5, "テスト管理者", true);

INSERT INTO quizzes (id, name, type) VALUES (1, "Slack", "macOS"), (2, "VSCode", "macOS"), (3, "Chrome", "macOS"), (4, "Github", "macOS");

INSERT INTO rankings (quiz_id, user_id, ranking) VALUES (1, 1, 1), (2, 2, 1), (3, 3, 1), (4, 4, 1);

INSERT INTO questions (id, quiz_id, contents) VALUES (1, 1, "DMを閲覧するには？");

INSERT INTO answers (id, question_id, contents, is_correct) VALUES (1, 1, "⌘+shit+↓", false), (2, 1, "⌘+shit+K", true), (3, 1, "⌘+K", false), (4, 1, "⌘+T", false);
