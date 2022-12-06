CREATE TABLE public.grades (
    id integer NOT NULL,
    hash text NOT NULL,
    gitlog text NOT NULL,
    openapi integer NOT NULL,
    openapilog text NOT NULL,
    golang integer NOT NULL,
    golanglog text NOT NULL,
    vue integer NOT NULL,
    vuelog text NOT NULL,
    docker integer NOT NULL,
    dockerlog text NOT NULL,
    lastcheck timestamp with time zone NOT NULL
);

CREATE TABLE public.students (
    id integer NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    repo_url text NOT NULL,
    public_key text NOT NULL,
    private_key text NOT NULL
);

ALTER TABLE ONLY public.grades
    ADD CONSTRAINT grades_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_email_key UNIQUE (email);

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_repo_url_key UNIQUE (repo_url);

ALTER TABLE ONLY public.grades
    ADD CONSTRAINT grades_id_fkey FOREIGN KEY (id) REFERENCES public.students(id) ON UPDATE CASCADE ON DELETE CASCADE;
