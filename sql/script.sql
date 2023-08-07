--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: clocks; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.clocks (
    id bigint NOT NULL,
    hour smallint NOT NULL,
    minute smallint NOT NULL,
    angle smallint NOT NULL,
    date timestamp with time zone
);


ALTER TABLE public.clocks OWNER TO root;

--
-- Name: clocks_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.clocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.clocks_id_seq OWNER TO root;

--
-- Name: clocks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.clocks_id_seq OWNED BY public.clocks.id;


--
-- Name: clocks id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.clocks ALTER COLUMN id SET DEFAULT nextval('public.clocks_id_seq'::regclass);


--
-- Data for Name: clocks; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.clocks (id, hour, minute, angle, date) FROM stdin;
1	1	0	30	2023-08-07 07:03:43.075065+00
2	2	0	60	2023-08-07 07:03:43.089467+00
3	3	0	90	2023-08-07 07:03:43.103399+00
4	4	0	120	2023-08-07 07:03:43.116225+00
5	5	0	150	2023-08-07 07:03:43.131971+00
6	6	0	180	2023-08-07 07:03:43.14278+00
7	7	0	150	2023-08-07 07:03:43.153024+00
8	8	0	120	2023-08-07 07:03:43.167218+00
9	9	0	90	2023-08-07 07:03:43.179889+00
10	10	0	60	2023-08-07 07:03:43.190795+00
11	11	0	30	2023-08-07 07:03:43.21602+00
12	12	0	0	2023-08-07 07:03:43.226964+00
\.


--
-- Name: clocks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.clocks_id_seq', 12, true);


--
-- Name: clocks clocks_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.clocks
    ADD CONSTRAINT clocks_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

-- Adminer 4.8.1 PostgreSQL 15.3 dump

\connect "pointersangle";

DROP TABLE IF EXISTS "clocks";
DROP SEQUENCE IF EXISTS clocks_id_seq;
CREATE SEQUENCE clocks_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."clocks" (
    "id" bigint DEFAULT nextval('clocks_id_seq') NOT NULL,
    "hour" smallint NOT NULL,
    "minute" smallint NOT NULL,
    "angle" smallint NOT NULL,
    "date" timestamptz,
    CONSTRAINT "clocks_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "clocks" ("id", "hour", "minute", "angle", "date") VALUES
(1,  1,  0, 30,  '2023-08-07 07:03:43.075065+00'),
(2,  2,  0, 60,  '2023-08-07 07:03:43.089467+00'),
(3,  3,  0, 90,  '2023-08-07 07:03:43.103399+00'),
(4,  4,  0, 120, '2023-08-07 07:03:43.116225+00'),
(5,  5,  0, 150, '2023-08-07 07:03:43.131971+00'),
(6,  6,  0, 180, '2023-08-07 07:03:43.14278+00'),
(7,  7,  0, 150, '2023-08-07 07:03:43.153024+00'),
(8,  8,  0, 120, '2023-08-07 07:03:43.167218+00'),
(9,  9,  0, 90,  '2023-08-07 07:03:43.179889+00'),
(10, 10, 0, 60,  '2023-08-07 07:03:43.190795+00'),
(11, 11, 0, 30,  '2023-08-07 07:03:43.21602+00'),
(12, 12, 0, 0,   '2023-08-07 07:03:43.226964+00');

-- 2023-08-07 07:04:45.396556+00
