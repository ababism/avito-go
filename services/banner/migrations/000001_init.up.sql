CREATE TABLE banner
(
    id         SERIAL PRIMARY KEY,
    content    JSONB     NOT NULL,
    is_active  BOOLEAN   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tag
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
--     updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE feature
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
--     updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE banners_tags
(
    banner_id INT NOT NULL,
    tag_id    INT NOT NULL,

    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banner (id),
    FOREIGN KEY (banner_id) REFERENCES tag (id)
);

CREATE TABLE banners_feature
(
    banner_id  INT NOT NULL,
    feature_id INT NOT NULL,

    PRIMARY KEY (banner_id, feature_id),
    FOREIGN KEY (banner_id) REFERENCES banner (id),
    FOREIGN KEY (banner_id) REFERENCES feature (id)
);

CREATE
    OR REPLACE FUNCTION update_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = NOW();
    RETURN NEW;
END;
$$
    language 'plpgsql';


CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE
    ON banner
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

