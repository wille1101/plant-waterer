SET client_encoding = 'UTF8';

-- Creation of users table
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY, 
    user_name VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL
);

-- Creation of plants table
CREATE TABLE IF NOT EXISTS plants (
    plant_id SERIAL PRIMARY KEY, 
    owner_id UUID NOT NULL,
    name VARCHAR(40),
    latin_name VARCHAR(40) DEFAULT 'Unknown latin name',
    last_watered DATE DEFAULT NOW(),
    watering_interval INT DEFAULT 7,
    water_within DATE,
    CONSTRAINT fk_owner
        FOREIGN KEY(owner_id)
            REFERENCES users(user_id)
            ON DELETE CASCADE
);

-- Creation of waterings table
CREATE TABLE IF NOT EXISTS waterings (
    timestamp DATE,
    plant_id INT,
    owner_id UUID,
    CONSTRAINT fk_plant
        FOREIGN KEY(plant_id)
            REFERENCES plants(plant_id)
            ON DELETE CASCADE,
    CONSTRAINT fk_owner
        FOREIGN KEY(owner_id)
            REFERENCES users(user_id)
            ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION log_watering()
  RETURNS TRIGGER 
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
	IF NEW.last_watered <> OLD.last_watered THEN
		 INSERT INTO waterings(timestamp, plant_id, owner_id)
		 VALUES(now(), OLD.plant_id, OLD.owner_id);
	END IF;

	RETURN NEW;
END;
$$;

CREATE OR REPLACE FUNCTION log_on_insert_watering()
  RETURNS TRIGGER 
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
	INSERT INTO waterings(timestamp, plant_id, owner_id)
	VALUES(now(), NEW.plant_id, NEW.owner_id);
	RETURN NEW;
END;
$$;

CREATE OR REPLACE FUNCTION change_water_within()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
        IF NEW.last_watered <> OLD.last_watered OR NEW.watering_interval <> OLD.watering_interval THEN
         UPDATE plants SET water_within = (last_watered + watering_interval)
         WHERE plants.plant_id = OLD.plant_id;
        END IF;

        RETURN NEW;
END;
$$;

CREATE OR REPLACE FUNCTION change_on_insert_water_within()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
        NEW.water_within = (NEW.last_watered + NEW.watering_interval);
        RETURN NEW;
END;
$$;

CREATE TRIGGER log_last_watered_changes
  BEFORE UPDATE
  ON plants
  FOR EACH ROW
  EXECUTE PROCEDURE log_watering();

CREATE TRIGGER log_on_insert_watered_changes
  AFTER INSERT 
  ON plants
  FOR EACH ROW
  EXECUTE PROCEDURE log_on_insert_watering();

CREATE TRIGGER update_water_within
  AFTER UPDATE
  ON plants
  FOR EACH ROW
  EXECUTE PROCEDURE change_water_within();

CREATE TRIGGER update_on_insert_water_within
  BEFORE INSERT
  ON plants
  FOR EACH ROW
  EXECUTE PROCEDURE change_on_insert_water_within();
