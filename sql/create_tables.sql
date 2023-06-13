-- Table: public.abilities

-- DROP TABLE IF EXISTS public.abilities;

CREATE TABLE
    IF NOT EXISTS public.abilities (
        id bigint NOT NULL DEFAULT nextval('abilities_id_seq':: regclass),
        ability text COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT abilities_pkey PRIMARY KEY (id)
    ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.abilities OWNER to postgres;

-- Table: public.universities

-- DROP TABLE IF EXISTS public.universities;

CREATE TABLE
    IF NOT EXISTS public.universities (
        id bigint NOT NULL DEFAULT nextval(
            'universities_id_seq':: regclass
        ),
        university text COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT universities_pkey PRIMARY KEY (id)
    ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.universities OWNER to postgres;

-- Table: public.personals

-- DROP TABLE IF EXISTS public.personals;

CREATE TABLE
    IF NOT EXISTS public.personals (
        id bigint NOT NULL DEFAULT nextval('personals_id_seq':: regclass),
        name text COLLATE pg_catalog."default" NOT NULL,
        surname text COLLATE pg_catalog."default" NOT NULL,
        username text COLLATE pg_catalog."default" NOT NULL,
        email text COLLATE pg_catalog."default" NOT NULL,
        password text COLLATE pg_catalog."default" NOT NULL,
        usertype text COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT personals_pkey PRIMARY KEY (id),
        CONSTRAINT personals_email_key UNIQUE (email),
        CONSTRAINT personals_username_key UNIQUE (username)
    ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.personals OWNER to postgres;

-- Table: public.personal_ability

-- DROP TABLE IF EXISTS public.personal_ability;

CREATE TABLE
    IF NOT EXISTS public.personal_ability (
        ability_id bigint NOT NULL,
        personal_id bigint NOT NULL,
        CONSTRAINT personal_ability_pkey PRIMARY KEY (ability_id, personal_id),
        CONSTRAINT fk_personal_ability_ability FOREIGN KEY (ability_id) REFERENCES public.abilities (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
        CONSTRAINT fk_personal_ability_personal FOREIGN KEY (personal_id) REFERENCES public.personals (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
    ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.personal_ability OWNER to postgres;

-- Table: public.personal_university

-- DROP TABLE IF EXISTS public.personal_university;

CREATE TABLE
    IF NOT EXISTS public.personal_university (
        university_id bigint NOT NULL,
        personal_id bigint NOT NULL,
        CONSTRAINT personal_university_pkey PRIMARY KEY (university_id, personal_id),
        CONSTRAINT fk_personal_university_personal FOREIGN KEY (personal_id) REFERENCES public.personals (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION,
        CONSTRAINT fk_personal_university_university FOREIGN KEY (university_id) REFERENCES public.universities (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
    ) TABLESPACE pg_default;

ALTER TABLE
    IF EXISTS public.personal_university OWNER to postgres;


 -- Table: public.experiences

-- DROP TABLE IF EXISTS public.experiences;

CREATE TABLE
    IF NOT EXISTS public.experiences (
        id bigint NOT NULL,
        company text COLLATE pg_catalog."default" NOT NULL,
        "position" text COLLATE pg_catalog."default",
        start_year date,
        finish_year date,
        personal_id bigint,
        CONSTRAINT experiences_pkey PRIMARY KEY (id),
        CONSTRAINT fk_personal_id FOREIGN KEY (personal_id) REFERENCES public.personals (id) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION NOT VALID
    ) TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.experiences OWNER to postgres;

-- Index: fki_fk_personal_id

-- DROP INDEX IF EXISTS public.fki_fk_personal_id;

CREATE INDEX
    IF NOT EXISTS fki_fk_personal_id ON public.experiences USING btree (personal_id ASC NULLS LAST) TABLESPACE pg_default;   