-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       balance INT NOT NULL DEFAULT 0
);

CREATE TABLE quests (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        cost INT NOT NULL
);

CREATE TABLE user_quests (
    user_id INT  NOT NULL,
    quest_id INT NOT NULL,
    CONSTRAINT "PK_64e5e3fba53c904dac4247759e6" PRIMARY KEY ("user_id", "quest_id")
);

ALTER TABLE "user_quests" ADD CONSTRAINT "FK_48036b74bf2716310b18d2d9468" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "user_quests" ADD CONSTRAINT "FK_48036b74bf2716310b18d2d9469" FOREIGN KEY ("quest_id") REFERENCES "quests"("id") ON DELETE CASCADE ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE "user_quests" DROP CONSTRAINT "FK_48036b74bf2716310b18d2d9468";
ALTER TABLE "user_quests" DROP CONSTRAINT "FK_48036b74bf2716310b18d2d9469";
DROP TABLE "user_quests";
DROP TABLE "quests";
DROP TABLE "users";
-- +goose StatementEnd
