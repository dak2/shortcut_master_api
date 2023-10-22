CREATE DATABASE IF NOT EXISTS shortcut_master_db;

USE shortcut_master_db;

CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name VARCHAR(256) NOT NULL,
  google_user_id VARCHAR(256) NOT NULL UNIQUE,
  email VARCHAR(256) NOT NULL UNIQUE,
  is_admin BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS quizzes (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name VARCHAR(256) NOT NULL,
  type ENUM("macOS", "windows"),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS rankings (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  quiz_id INT NOT NULL,
  user_id INT NOT NULL,
  ranking INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS questions (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  quiz_id INT NOT NULL,
  quiz_type VARCHAR(256) NOT NULL,
  contents VARCHAR(256) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS answers (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  question_id INT NOT NULL,
  contents VARCHAR(256) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (question_id) REFERENCES questions(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS answer_histories (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  answer_id INT NOT NULL,
  contents VARCHAR(256) NOT NULL,
  is_correct BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (answer_id) REFERENCES answers(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8;

INSERT INTO
  users (id, name, google_user_id, email, is_admin)
VALUES
  (1, "テストユーザー1", 1, "test1@example.com", false),
  (2, "テストユーザー2", 2, "test2@example.com", false),
  (3, "テストユーザー3", 3, "test3@example.com", false),
  (4, "テストユーザー4", 4, "test4@example.com", false),
  (5, "テスト管理者", 5, "test5@example.com", true);

INSERT INTO
  quizzes (id, name, type)
VALUES
  (1, "Slack", "macOS"),
  (2, "VSCode", "macOS"),
  (3, "Chrome", "macOS"),
  (4, "GitHub", "macOS");

INSERT INTO
  questions (id, quiz_id, quiz_type, contents)
VALUES
  (1, 1, 'slack', "メッセージ送信の取り消し"),
  (2, 1, 'slack', "検索を開始する"),
  (3, 1, 'slack', "別の会話へ移動する"),
  (4, 1, 'slack', "履歴で前に戻る"),
  (5, 1, 'slack', "DMを閲覧する"),
  (6, 1, 'slack', "「後で」のアイテムを表示する"),
  (7, 1, 'slack', "チャンネルをブラウズする"),
  (8, 1, 'slack', "メッセージを未読にする"),
  (9, 1, 'slack', "全未読画面を開く"),
  (10, 1, 'slack', "スレッド画面を開く");

INSERT INTO
  answers (id, question_id, contents)
VALUES
  (1, 1, "⌘+Z"),
  (2, 2, "⌘+G"),
  (3, 3, "⌘+K"),
  (4, 4, "⌘+["),
  (5, 5, "⌘+Shift+K"),
  (6, 6, "⌘+Shift+S"),
  (7, 7, "⌘+Shift+L"),
  (8, 8, "Option+Click"),
  (9, 9, "⌘+Shift+A"),
  (10, 10, "⌘+Shift+T");

-- INSERT INTO answer_histories (id, answer_id, contents, is_correct) VALUES (1, 1, "⌘+Z", true), (2, 2, "⌘+G", false), (3, 3, "⌘+K", false), (4, 4, "⌘+K", false), (5, 5, "⌘+Shift+K", false), (6, 6, "⌘+Shift+S", false), (7, 7, "⌘+Shift+L", false), (8, 8, "Option+Click", false), (9, 9, "⌘+Shift+A", false), (10, 10, "⌘+Shift+T", false);
INSERT INTO
  rankings (quiz_id, user_id, ranking)
VALUES
  (1, 1, 1),
  (2, 2, 1),
  (3, 3, 1),
  (4, 4, 1);
