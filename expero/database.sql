CREATE TABLE IF NOT EXISTS person_discounts
(
    id bigint NOT NULL
    PRIMARY KEY,
    person_id bigint NOT NULL
    CONSTRAINT fk_person_discounts_person_id REFERENCES person,
    pizzeria_id bigint NOT NULL
    CONSTRAINT fk_person_discounts_pizzeria_id REFERENCES pizzeria,
    discount numeric DEFAULT 0 NOT NULL
);