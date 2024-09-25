CREATE TABLE public.admins (
    id bigint NOT NULL,
    username character varying(64) NOT NULL,
    password text NOT NULL
);

ALTER TABLE public.admins ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.admins_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.customers (
    id bigint NOT NULL,
    language_id bigint NOT NULL,
    phone_humber character varying(20) NOT NULL
);

ALTER TABLE public.customers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.languages (
    id bigint NOT NULL,
    short_name character(2) NOT NULL,
    full_name character varying(32) NOT NULL
);

ALTER TABLE public.languages ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.languages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.order_statuses (
    id bigint NOT NULL,
    ready_at timestamp without time zone,
    returned_at timestamp without time zone,
    customer_notified_at timestamp without time zone,
    is_outsourced boolean DEFAULT false NOT NULL,
    is_receipt_lost boolean DEFAULT false NOT NULL
);

ALTER TABLE public.order_statuses ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.order_statuses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.order_types (
    id bigint NOT NULL,
    short_name character varying(16) NOT NULL,
    full_name character varying(32) NOT NULL
);

ALTER TABLE public.order_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.order_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.orders (
    id bigint NOT NULL,
    order_status_id bigint NOT NULL,
    order_type_id bigint NOT NULL,
    worker_id bigint NOT NULL,
    customer_id bigint NOT NULL,
    reason text,
    defect text,
    total_price character varying(64),
    prepayment character varying(64)
);

ALTER TABLE public.orders ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.wokerks (
    id bigint NOT NULL,
    first_name character varying(64) NOT NULL,
    last_name character varying(64) NOT NULL
);

ALTER TABLE public.wokerks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.wokerks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_short_name_unique UNIQUE (short_name);

ALTER TABLE ONLY public.order_statuses
    ADD CONSTRAINT order_statuses_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.order_types
    ADD CONSTRAINT order_types_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.order_types
    ADD CONSTRAINT order_types_short_name_unique UNIQUE (short_name);

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.wokerks
    ADD CONSTRAINT wokerks_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customer_language_fk FOREIGN KEY (language_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_customer_id_fk FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_order_status_id_fk FOREIGN KEY (order_status_id) REFERENCES public.order_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_order_type_id_fk FOREIGN KEY (order_type_id) REFERENCES public.order_types(id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_worker_id_fk FOREIGN KEY (worker_id) REFERENCES public.wokerks(id) ON UPDATE CASCADE ON DELETE CASCADE;
