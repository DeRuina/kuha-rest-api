--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0 (Debian 17.0-1.pgdg110+1)
-- Dumped by pg_dump version 18.0 (Ubuntu 18.0-1.pgdg24.04+3)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: injuries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.injuries (
    competitor_id integer NOT NULL,
    injury_type integer NOT NULL,
    severity integer,
    pain_level integer,
    description character varying(255),
    date_start timestamp without time zone NOT NULL DEFAULT NOW(),
    status integer NOT NULL DEFAULT 0,
    date_end timestamp without time zone,
    injury_id integer,
    meta character varying(255)
);


--
-- Name: querys; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.querys (
    competitor_id integer NOT NULL,
    query_type integer,
    answers character varying(255),
    comment character varying(255),
    "timestamp" timestamp without time zone NOT NULL,
    meta character varying(255)
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    username integer,
    role_id integer,
    password character varying(255),
    email character varying(255),
    share_permission integer,
    collect_permission integer,
    eula_accepted integer,
    create_time timestamp without time zone,
    message character varying(255),
    salt character varying(255)
);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON SCHEMA public TO kamkadmin;


--
-- PostgreSQL database dump complete
--

